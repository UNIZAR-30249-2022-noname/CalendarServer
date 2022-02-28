package uploaddatarepositorymysql_test

import (
	"testing"

	uploaddatarepositorymysql "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/MySQL/UploadDataRepository"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	consultas "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/sql"
	"github.com/stretchr/testify/assert"
)

//Creates a degree (OK) and then creates the same degree (FAIL)
func TestCreateDegree(t *testing.T) {
	t.Skip() //remove for activating it
	//Prepare
	assert := assert.New(t)
	repos := uploaddatarepositorymysql.New()
	//Test
	res, err := repos.CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto")
	assert.Equal(err, nil, "There shouldn't be an error")
	assert.Equal(res, true, "Should be true")
	res, err = repos.CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto")
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	//Delete
	repos.RawExec(consultas.TruncDegree)
}

//Creates a subject (OK) and then creates the same subject (FAIL)
//Creates a subject but the degree isn't in the database (FAIL)
func TestCreateSubject(t *testing.T) {
	t.Skip() //remove for activating it
	//Prepare
	assert := assert.New(t)
	repos := uploaddatarepositorymysql.New()
	repos.CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto")

	//Test
	res, err := repos.CreateNewSubject(25802, "Informática", 558)
	assert.Equal(err, nil, "There shouldn't be an error")
	assert.Equal(res, true, "Should be true")
	res, err = repos.CreateNewSubject(25802, "Informática", 558)
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	res, err = repos.CreateNewSubject(25803, "NoExisteSuTitulaciónAsíQueFalla", 560)
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	//Delete
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncDegree)
}

//Creates a year (OK) and then creates the same year (FAIL)
//Creates a year but the degree isn't in the database (FAIL)
func TestCreateYear(t *testing.T) {
	t.Skip() //remove for activating it
	//Prepare
	assert := assert.New(t)
	repos := uploaddatarepositorymysql.New()
	repos.CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto")

	//Test
	res, err := repos.CreateNewYear(1, 558)
	assert.Equal(err, nil, "There shouldn't be an error")
	assert.Equal(res, true, "Should be true")
	res, err = repos.CreateNewYear(1, 558)
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	res, err = repos.CreateNewYear(2, 560)
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	//Delete
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncYear)
}

//Creates a group (OK) and then creates the same group (FAIL)
//Creates a group but the year isn't in the database (FAIL)
func TestCreateGroup(t *testing.T) {
	t.Skip() //remove for activating it
	//Prepare
	assert := assert.New(t)
	repos := uploaddatarepositorymysql.New()
	repos.CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto")
	repos.CreateNewYear(1, 558) //Id will be 5581
	//Test
	res, err := repos.CreateNewGroup(1, 5581)
	assert.Equal(err, nil, "There shouldn't be an error")
	assert.Equal(res, true, "Should be true")
	res, err = repos.CreateNewGroup(1, 5581)
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	res, err = repos.CreateNewGroup(1, 5582)
	assert.Equal(err, apperrors.ErrSql, "There should be an error")
	assert.Equal(res, false, "Should be false")
	//Delete
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncGroup)
}

//Creates hour (See the comments)
func TestCreateHour(t *testing.T) {
	t.Skip() //remove for activating it
	//Prepare
	assert := assert.New(t)
	repos := uploaddatarepositorymysql.New()
	repos.CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto")
	repos.CreateNewYear(1, 558) //Id will be 5581
	repos.CreateNewGroup(1, 5581)
	repos.CreateNewSubject(25802, "Informática", 558)
	//Test
	//OK Test
	res, err := repos.CreateNewHour(3500, 3500, 25802, 55811, constants.THEORICAL, "", "")
	assert.Equal(err, nil, "There shouldn't be an error")
	assert.Equal(res, true, "Should be true")
	//Hour kind=PRACTICES without week (FAIL)
	res, err = repos.CreateNewHour(3500, 3500, 25802, 55811, constants.PRACTICES, "", "")
	assert.Equal(err, apperrors.ErrInvalidKind, "There should be an invalid kind error")
	assert.Equal(res, false, "Should be false")
	//Hour kind=PRACTICES without group (FAIL)
	res, err = repos.CreateNewHour(3500, 3500, 25802, 55811, constants.PRACTICES, "", "a")
	assert.Equal(err, apperrors.ErrInvalidKind, "There should be an invalid kind error")
	assert.Equal(res, false, "Should be false")
	//Hour kind=EXERCISES without group (FAIL)
	res, err = repos.CreateNewHour(3500, 3500, 25802, 55811, constants.EXERCISES, "", "")
	assert.Equal(err, apperrors.ErrInvalidKind, "There should be an invalid kind error")
	assert.Equal(res, false, "Should be false")
	//Hour kind that not exists (FAIL)
	res, err = repos.CreateNewHour(3500, 3500, 25802, 55811, 4, "", "")
	assert.Equal(err, apperrors.ErrInvalidKind, "There should be an invalid kind error")
	assert.Equal(res, false, "Should be false")
	//Hour invalid subjectCode (FAIL)
	res, err = repos.CreateNewHour(3500, 3500, 25803, 55811, constants.THEORICAL, "", "")
	assert.Equal(err, apperrors.ErrSql, "There should be an sql error")
	assert.Equal(res, false, "Should be false")
	//Hour invalid groupCode (FAIL)
	res, err = repos.CreateNewHour(3500, 3500, 25802, 55812, constants.THEORICAL, "", "")
	assert.Equal(err, apperrors.ErrSql, "There should be an sql error")
	assert.Equal(res, false, "Should be false")
	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
}
