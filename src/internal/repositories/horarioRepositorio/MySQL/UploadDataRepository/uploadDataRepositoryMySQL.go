package uploaddatarepositorymysql

import (
	"database/sql"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	consultas "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/sql"
	_ "github.com/go-sql-driver/mysql"
)

type HorarioRepositorioMySQL struct {
	db *sql.DB
}

func New() *HorarioRepositorioMySQL {
	db, _ := sql.Open("mysql", "user:user@tcp(127.0.0.1:6033)/app_db")
	return &HorarioRepositorioMySQL{db}
}

func (repo *HorarioRepositorioMySQL) CloseConn() error {
	return repo.db.Close()
}

func (repo *HorarioRepositorioMySQL) RawExec(exec string) error {
	_, err := repo.db.Exec(exec)
	return err
}

//EntryFound is a function which returns true if the given
//entry [Entry] is in the database
func (repo *HorarioRepositorioMySQL) EntryFound(entry domain.Entry) bool {

	res, err := repo.db.Query(consultas.SearchEntry,
		domain.HourToInt(entry.Init), domain.HourToInt(entry.End),
		entry.Subject.Kind, entry.Week, entry.Group, entry.Subject.Name)
	found := res.Next()
	_ = err
	return found
}

func (repo *HorarioRepositorioMySQL) CreateNewDegree(id int, name string) (bool, error) {
	//Create degree given an id and a name
	_, err := repo.db.Query(consultas.CreateDegree, id, name)
	if err != nil {
		return false, apperrors.ErrSql
	}
	return true, nil
}

func (repo *HorarioRepositorioMySQL) CreateNewSubject(id int, name string, degreeCode int) (bool, error) {
	//Create a subject given an id and a name and the degreeCode
	_, err := repo.db.Query(consultas.CreateSubject, id, id, name, degreeCode)
	if err != nil {
		return false, apperrors.ErrSql
	}
	return true, nil
}

func (repo *HorarioRepositorioMySQL) CreateNewYear(year int, degreeCode int) (bool, error) {
	//Create a subject given an id and a name and the degreeCode
	id := degreeCode*10 + year
	_, err := repo.db.Query(consultas.CreateYear, id, year, degreeCode)
	if err != nil {
		return false, apperrors.ErrSql
	}
	return true, nil
}

func (repo *HorarioRepositorioMySQL) CreateNewGroup(group int, yearCode int) (bool, error) {
	//Create a subject given an id and a name and the degreeCode
	id := yearCode*10 + group
	_, err := repo.db.Query(consultas.CreateGroup, id, group, yearCode)
	if err != nil {
		return false, apperrors.ErrSql
	}
	return true, nil
}

func (repo *HorarioRepositorioMySQL) CreateNewHour(available, total, subjectCode, groupCode, kind int, group, week string) (bool, error) {
	//Cuidado que las horas tipo 2 son clases de problemas y las tipo 3 prácticas
	//Create a subject given an id and a name and the degreeCode
	if kind == constants.PRACTICES {
		if week == "" || &week == nil {
			return false, apperrors.ErrInvalidKind
		}
		if group == "" || &group == nil {
			return false, apperrors.ErrInvalidKind
		}
	} else if kind == constants.EXERCISES {
		if group == "" || &group == nil {
			return false, apperrors.ErrInvalidKind
		}
	} else if kind < 1 || kind > 3 {
		return false, apperrors.ErrInvalidKind
	}
	_, err := repo.db.Query(consultas.CreateHour, available, total, kind, group, week, subjectCode, groupCode)
	if err != nil {
		return false, apperrors.ErrSql
	}
	return true, nil
}
