package horarioRepositorio

import (
	"database/sql"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

func New() *repo {
	db, _ := sql.Open("mysql", "user:user@tcp(127.0.0.1:6033)/app_db")
	return &repo{db}
}

type AuxAvaiableHours struct {
	Kind      int    `json:"tipo"`
	Subject   string `json:"name"`
	Remaining int    `json:"disponibles"`
	Max       int    `json:"totales"`
}

func (repo *repo) CloseConn() (error) {
	return repo.db.Close();
}

func (repo *repo) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {

	res := make([]domain.AvailableHours, 0)
	results, err := repo.db.Query("SELECT a.disponibles, a.totales, a.tipo, a.nombre FROM "+
		"(SELECT hora.*, a.nombre FROM hora INNER JOIN "+
		"(SELECT * FROM asignatura "+
		"WHERE id IN "+
		"(SELECT perteneceA.idA FROM perteneceA "+
		"INNER JOIN titulacion ON titulacion.id=perteneceA.idT WHERE titulacion.nombre=?)) a ON a.id=hora.idasignatura) a "+
		"INNER JOIN "+
		"(SELECT * FROM hora WHERE id "+
		"IN (SELECT grupodocente.id FROM `grupodocente` "+
		"INNER JOIN curso ON grupodocente.idcurso=curso.id WHERE curso.numero=? AND grupodocente.numero=?)) b "+
		"ON a.id=b.id", terna.Titulacion, terna.Curso, terna.Grupo)
	if err != nil {
		return []domain.AvailableHours{}, apperrors.ErrSql
	}
	
	for results.Next() {
		var auxv AuxAvaiableHours
		// for each row, scan the result into our tag composite object
		err = results.Scan(&auxv.Remaining, &auxv.Max, &auxv.Kind, &auxv.Subject)
		if err != nil {
			return []domain.AvailableHours{}, err
		}
		res = append(res, domain.AvailableHours{Kind: auxv.Kind, Subject: auxv.Subject, Remaining: auxv.Remaining, Max: auxv.Max})
	}

	return res, nil
}
