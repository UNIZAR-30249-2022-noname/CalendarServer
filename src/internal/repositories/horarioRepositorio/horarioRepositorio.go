package horarioRepositorio

import (
	"database/sql"
	"log"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

func New() *repo {
	db, _ := sql.Open("mysql", "user:user@tcp(127.0.0.1:6033)/horarios_db")
	defer db.Close()
	return &repo{db}
}

type AuxAvaiableHours struct {
	Kind      int    `json:"tipo"`
	Subject   string `json:"name"`
	Remaining int    `json:"disponibles"`
	Max       int    `json:"totales"`
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
		return []domain.AvailableHours{}, err
	}
	i := 0
	for results.Next() {
		var auxv AuxAvaiableHours
		// for each row, scan the result into our tag composite object
		err = results.Scan(&auxv.Remaining, &auxv.Max, &auxv.Kind, &auxv.Subject)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		res[i] = domain.AvailableHours{Kind: auxv.Kind, Subject: auxv.Subject, Remaining: auxv.Remaining, Max: auxv.Max}
		log.Printf("%v", auxv)
	}

	return res, nil
}

/*
func TestIsSuperAnimal(t *testing.T) {

	AvailableHours := []domain.AvailableHours{
		{
			Kind:      domain.TEORIA,
			Subject:   "IC",
			Remaining: 5,
			Max:       5,
		},
	}

	repositorio := horarioRepositorio.New()
	expected := true
	got := true
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
*/
