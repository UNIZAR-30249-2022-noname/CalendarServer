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

func (repo *repo) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {

	res := make([]domain.AvailableHours, 0)
	results, err := repo.db.Query(consultas.SelectAvaiableHours, terna.Titulacion, terna.Curso, terna.Grupo)
	if err != nil {
		return []domain.AvailableHours{}, apperrors.ErrSql
	}
	
	for results.Next() {
		var auxv models.AuxAvaiableHours
		// for each row, scan the result into our tag composite object
		err = results.Scan(&auxv.Remaining, &auxv.Max, &auxv.Kind, &auxv.Subject)
		if err != nil {
			return []domain.AvailableHours{}, apperrors.ErrSql
		}
		res = append(res, models.AuxToReal(auxv))
	}

	return res, nil
}

func (repo *repo) CreateNewEntry(entry domain.Entry) (error) {
	var idhoras, idgrupo, idaula int
	err := repo.db.QueryRow(consultas.SelectIdHoraGrupo,
			entry.Subject.Kind, entry.Group, entry.Week, entry.Subject.Name).Scan(&idhoras, &idgrupo)
	if err != nil { return apperrors.ErrSql }
	err = repo.db.QueryRow(consultas.SelectIdAula, entry.Room.Name).Scan(&idaula)
	if err != nil { return apperrors.ErrSql }
	now := time.Now()
  	ultModificacion := now.Format("2006-02-01")
	res, err := repo.db.Exec(consultas.InsertEntradaHorario, domain.HourToInt(entry.Init), 						 
				domain.HourToInt(entry.End), idhoras, idaula, idgrupo, ultModificacion)
	if err != nil { return apperrors.ErrSql }
	count, err := res.RowsAffected()
	if err != nil || count < 1 { return apperrors.ErrSql }
	repo.updateHours(entry.Init,entry.End,idhoras,true)
	return nil
}

func (repo *repo) DeleteEntry(entry domain.Entry) (error){
	var idhoras, idgrupo, idaula int
	err := repo.db.QueryRow(consultas.SelectIdHoraGrupo,
		entry.Subject.Kind, entry.Group, entry.Week, entry.Subject.Name).Scan(&idhoras, &idgrupo)
	if err != nil { return apperrors.ErrSql }
	err = repo.db.QueryRow(consultas.SelectIdAula, entry.Room.Name).Scan(&idaula)
	if err != nil { return apperrors.ErrSql }
	res , err := repo.db.Exec(consultas.DeleteEntradaHorario, domain.HourToInt(entry.Init), 						 
	domain.HourToInt(entry.End), idhoras, idaula, idgrupo)
	if err != nil { return apperrors.ErrSql }
	count, err := res.RowsAffected()
	if err != nil || count < 1 { return apperrors.ErrSql }
	repo.updateHours(entry.Init,entry.End,idhoras,false)
	return nil
}

func (repo *repo) updateHours(ini, fin domain.Hour, idhora int, create bool) (error){
	//Create es para quitar las horas de disponibles si es true y añadirlas si es false
	var horastotales, horasdisponibles, newhDisponibles int
	err := repo.db.QueryRow(consultas.SearchHours,
		idhora).Scan(&horastotales, &horasdisponibles)
	if err != nil { return apperrors.ErrSql }
	hDisponibles := domain.IntToHour(horasdisponibles)
	haQuitar := domain.IntToHour(domain.SubstractHour(fin, ini))
	if create {
		newhDisponibles = domain.SubstractHour(hDisponibles,haQuitar)
		if newhDisponibles < 0 { return apperrors.ErrIllegalOperation }
	} else {
		newhDisponibles = domain.AddHour(hDisponibles,haQuitar)
		if newhDisponibles > horastotales { return apperrors.ErrIllegalOperation }
	}
	//TODO Updateo
	res , err := repo.db.Exec(consultas.UpdateHours, newhDisponibles, idhora)
	if err != nil { return apperrors.ErrSql }
	count, err := res.RowsAffected()
	if err != nil || count < 1 { return apperrors.ErrSql }
	return nil
}