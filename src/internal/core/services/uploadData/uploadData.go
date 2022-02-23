package uploaddata

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

type UploadDataserviceImp struct {
	horarioRepositorio ports.HorarioRepositorio
}

func New() *UploadDataserviceImp {
	return &UploadDataserviceImp{}
}

//3->Code subject, 4->Name subject, 11->Code Degree, 12->Name degree
//17->Year (1,2...), 23->nº groups, 30->Hours t1, 32->Hours t2, 34->Hours t3
//71->nº subgroups t1, 72->nº subgroups t1
func (srv *UploadDataserviceImp) UpdateByCSV(csv string) (bool, error) {
	lines := strings.Split(csv, "\n")
	groups := make([]int, 500, 500)
	subjects := make([]int, 1000, 1000)
	prevDegree := 0
	prevYear := 0
	subjectsIn := 0
	groupsIn := 0
	hoursIn := 0
	longstring := "INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES "
	longstringGroup := "INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES "
	longstringHours := "INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `idasignatura`, `idgrupo`, `grupo`, `semana`) VALUES "
	for i, actLine := range lines {
		if i < 3 || lines[i] == "" || len(lines[i]) < 72 {
			continue
		}
		cells := strings.Split(actLine, ";")
		subjectId, _ := strconv.Atoi(cells[3])
		subjectName := cells[4]
		fmt.Println("Asignatura " + strconv.Itoa(subjectId) + ": " + subjectName)
		degreeId, _ := strconv.Atoi(cells[11])
		degreeName := cells[12]
		fmt.Println("Grado " + strconv.Itoa(degreeId) + ": " + degreeName)
		year, _ := strconv.Atoi(cells[17])
		nGroups, _ := strconv.Atoi(cells[23])
		nGroupsT2, _ := strconv.Atoi(cells[71])
		nGroupsT2 /= 2
		nGroupsT3, _ := strconv.Atoi(cells[72])
		nGroupsT3 /= 2

		aux := strings.Split(cells[30], ",")
		hoursT1, _ := strconv.Atoi(aux[0])
		aux = strings.Split(cells[32], ",")
		hoursT2, _ := strconv.Atoi(aux[0])
		aux = strings.Split(cells[34], ",")
		hoursT3, _ := strconv.Atoi(aux[0])

		if prevDegree != degreeId {
			srv.horarioRepositorio.CreateNewDegree(degreeId, degreeName)
			fmt.Println("Nuevo grado")
			prevDegree = degreeId
		}

		actYearId := degreeId*10 + year
		if prevYear != actYearId {
			srv.horarioRepositorio.CreateNewYear(year, degreeId)
			fmt.Println("Nuevo año")
			prevYear = actYearId
		}

		if !contains(subjects, subjectId) {
			if subjectsIn > 0 {
				longstring = longstring + ", "
			}
			longstring = longstring + "('" + strconv.Itoa(subjectId) + "','" + strconv.Itoa(subjectId) + "','" + subjectName + "','" + strconv.Itoa(degreeId) + "')"
			subjectsIn++
			subjects = append(subjects, subjectId)

			for j := 1; j <= nGroups; j++ {
				actGroupId := actYearId*10 + j
				if !contains(groups, actGroupId) {
					if groupsIn > 0 {
						longstringGroup = longstringGroup + ", "
					}
					longstringGroup = longstringGroup + "('" + strconv.Itoa(actGroupId) + "','" + strconv.Itoa(j) + "','" + strconv.Itoa(actYearId) + "')"
					groupsIn++
					groups = append(groups, actGroupId)
				}

				if hoursIn > 0 {
					longstringHours += ", "
				}
				longstringHours = longstringHours + "(NULL,'" + strconv.Itoa(hoursT1*100) + "','" + strconv.Itoa(hoursT1*100) + "','" + strconv.Itoa(constants.THEORICAL) + "','" + strconv.Itoa(subjectId) + "','" + strconv.Itoa(actGroupId) + "','','')"
				hoursIn++
				for k := 1; k <= nGroupsT2; k++ {
					if hoursIn > 0 {
						longstringHours += ", "
					}
					longstringHours = longstringHours + "(NULL,'" + strconv.Itoa(hoursT2*100) + "','" + strconv.Itoa(hoursT2*100) + "','" + strconv.Itoa(constants.EXERCISES) + "','" + strconv.Itoa(subjectId) + "','" + strconv.Itoa(actGroupId) + "','" + strconv.Itoa(k) + "','')"
					hoursIn++
				}
				for k := 1; k <= nGroupsT3; k++ {
					if hoursIn > 0 {
						longstringHours += ", "
					}
					longstringHours = longstringHours + "(NULL,'" + strconv.Itoa(hoursT3*100) + "','" + strconv.Itoa(hoursT3*100) + "','" + strconv.Itoa(constants.PRACTICES) + "','" + strconv.Itoa(subjectId) + "','" + strconv.Itoa(actGroupId) + "','" + strconv.Itoa(k) + "','a')"
					hoursIn++
				}
			}
		}
	}
	fmt.Println(longstringGroup)
	err := srv.horarioRepositorio.RawExec(longstringGroup)
	if err != nil {
		fmt.Println("Fallo de longstring group" + err.Error())
		return false, err
	}
	if subjectsIn > 0 {
		fmt.Println(longstring)
		err := srv.horarioRepositorio.RawExec(longstring)
		if err != nil {
			fmt.Println("Fallo de longstring " + err.Error())
			return false, err
		}
	}
	if hoursIn > 0 {
		fmt.Println(longstringHours)
		err := srv.horarioRepositorio.RawExec(longstringHours)
		if err != nil {
			fmt.Println("Fallo de longstring " + err.Error())
			return false, err
		}
	}
	return true, nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
