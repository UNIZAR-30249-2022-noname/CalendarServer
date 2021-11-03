package horarioRepositorio

import (
	"database/sql"
	"fmt"
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
			entry.Subject.Kind, entry.Subject.Name).Scan(&idhoras, &idgrupo)
	if err != nil { return apperrors.ErrSql }
	err = repo.db.QueryRow(consultas.SelectIdAula, entry.Room.Name).Scan(&idaula)
	if err != nil { return apperrors.ErrSql }
	now := time.Now()
  	ultModificacion := now.Format("2006-02-01")
	results, err := repo.db.Exec(consultas.InsertEntradaHorario, entry.Init, entry.End, idhoras, idaula, idgrupo, ultModificacion)
	fmt.Printf("%v", results)
	if err != nil { return apperrors.ErrSql }
	return nil
}