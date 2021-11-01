package horarioRepositorio

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/mysql"
)

type repo struct {
	horarioRepositorio ports.HorarioRepositorio
}

func New(horarioRepositorio ports.HorarioRepositorio) *repo {
	return &repo{horarioRepositorio: horarioRepositorio}
}


func (srv *repo) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {
	db, err := godb.Open(mysql.Adapter, "./horario.db")
	res := make([]domain.AvailableHours,0)
	err = db.SelectFrom("books").
		Columns("author", "count(*) as count").
		GroupBy("author").
		Having("count(*) > 3").
		Do(&res)
	if err != nil {
		return []domain.AvailableHours{}, err
	}

	return res, nil
}