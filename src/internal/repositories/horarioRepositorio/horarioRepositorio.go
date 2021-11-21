package horarioRepositorio

import (
	"database/sql"
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/models"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	consultas "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/sql"
	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

func New() *repo {
	db, _ := sql.Open("mysql", "user:user@tcp(127.0.0.1:6033)/app_db")
	return &repo{db}
}

func (repo *repo) CloseConn() (error) {
	return repo.db.Close();
}

//GetAvaiabledHours is a function which returns a set of [AvailableHours]
//given a completed [Terna]
func (repo *repo) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {

	res := make([]domain.AvailableHours, 0)
	//Select hours given a degree, course and group 
	//(returns remaining hours, max hours, kind and the subject's name)
	results, err := repo.db.Query(consultas.SelectAvaiableHours, terna.Titulacion, terna.Curso, terna.Grupo)
	if err != nil {
		return []domain.AvailableHours{}, apperrors.ErrSql
	}
	
	for results.Next() {	//SQL iteration loop
		var auxv models.AuxAvaiableHours
		// for each row, scan the result into our tag composite object
		err = results.Scan(&auxv.Remaining, &auxv.Max, &auxv.Kind, &auxv.Subject)
		if err != nil {
			return []domain.AvailableHours{}, apperrors.ErrSql
		}
		res = append(res, models.AuxToReal(auxv)) //We introduce the result to the slice
	}

	return res, nil
}

//CreateNewEntry is a function which creates an entry in the database
//given a completed [Entry] 
func (repo *repo) CreateNewEntry(entry domain.Entry) (error) {
	var idhoras, idgrupo, idaula int
	//We get idhoras & idgrupo for the entry
	err := repo.db.QueryRow(consultas.SelectIdHoraGrupo,
			entry.Subject.Kind, entry.Group, entry.Week, entry.Subject.Name).Scan(&idhoras, &idgrupo)
	if err != nil { return apperrors.ErrSql }
	//We get idaula for the entry
	err = repo.db.QueryRow(consultas.SelectIdAula, entry.Room.Name).Scan(&idaula)
	if err != nil { return apperrors.ErrSql }
	now := time.Now()
  	ultModificacion := now.Format("2006-02-01")	//Sacamos la fecha formateada para introducirla
	//Insert
	res, err := repo.db.Exec(consultas.InsertEntradaHorario, domain.HourToInt(entry.Init), 						 
				domain.HourToInt(entry.End), idhoras, idaula, idgrupo, ultModificacion)
	if err != nil { return apperrors.ErrSql }
	count, err := res.RowsAffected() //We check if something has been changed just in case
	if err != nil || count < 1 { return apperrors.ErrSql }
	repo.updateHours(entry.Init,entry.End,idhoras,true)
	return nil
}

//CreateNewEntry is a function which deletes an entry in the database
//given a completed [Entry] 
func (repo *repo) DeleteEntry(entry domain.Entry) (error){
	var idhoras, idgrupo, idaula int
	//We get idhoras & idgrupo for the entry
	err := repo.db.QueryRow(consultas.SelectIdHoraGrupo,
		entry.Subject.Kind, entry.Group, entry.Week, entry.Subject.Name).Scan(&idhoras, &idgrupo)
	if err != nil { return apperrors.ErrSql }
	//We get idaula for the entry
	err = repo.db.QueryRow(consultas.SelectIdAula, entry.Room.Name).Scan(&idaula)
	if err != nil { return apperrors.ErrSql }
	//Delete
	res , err := repo.db.Exec(consultas.DeleteEntradaHorario, domain.HourToInt(entry.Init), 						 
	domain.HourToInt(entry.End), idhoras, idaula, idgrupo)
	if err != nil { return apperrors.ErrSql }
	count, err := res.RowsAffected() //We check if something has been changed just in case
	if err != nil || count < 1 { return apperrors.ErrSql }
	repo.updateHours(entry.Init,entry.End,idhoras,false)
	return nil
}

//updateHours is a function which updates the avaiable hours (hora.disponibles)
//in the database given an initial and final [Hour], an id of the hora row to update
//and a boolean (true if it was a create -> substracts the hour 
//				and false if it was a delete -> adds the hour)
func (repo *repo) updateHours(ini, fin domain.Hour, idhora int, create bool) (error){
	//Create is to remove the available hours if true and add them if false
	var horastotales, horasdisponibles, newhDisponibles int
	//We get the total and available hours from 'hora'
	err := repo.db.QueryRow(consultas.SearchHours,
		idhora).Scan(&horastotales, &horasdisponibles)
	if err != nil { return apperrors.ErrSql }
	hDisponibles := domain.IntToHour(horasdisponibles)
	//Hours to add or remove from the available
	haQuitar := domain.IntToHour(domain.SubstractHour(fin, ini))
	if create { //If it is CreateNewEntry
		newhDisponibles = domain.SubstractHour(hDisponibles,haQuitar) //Substract hours
		if newhDisponibles < 0 { return apperrors.ErrIllegalOperation }
	} else { //If it is DeleteEntry
		newhDisponibles = domain.AddHour(hDisponibles,haQuitar) //Add hours
		if newhDisponibles > horastotales { return apperrors.ErrIllegalOperation }
	}
	res , err := repo.db.Exec(consultas.UpdateHours, newhDisponibles, idhora)
	if err != nil { return apperrors.ErrSql }
	count, err := res.RowsAffected()
	if err != nil || count < 1 { return apperrors.ErrSql }
	return nil
}

func (repo *repo) RawExec(exec string) (error){
	_ , err := repo.db.Exec(exec)
	return err
}

//Sirve para ver si una entrada existe en la BD
func (repo *repo) EntryFound(entry domain.Entry) (bool){

	res, err := repo.db.Query(consultas.SearchEntry,
		domain.HourToInt(entry.Init), domain.HourToInt(entry.End), 
		entry.Subject.Kind, entry.Week, entry.Group, entry.Subject.Name)
	found := res.Next()
	_ = err
	return found
}