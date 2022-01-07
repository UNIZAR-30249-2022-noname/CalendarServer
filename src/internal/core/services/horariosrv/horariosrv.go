package horariosrv

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	ics "github.com/arran4/golang-ical"
)

//HorarioServiceImp is the implemetation of [HorarioService] interface.
type HorarioServiceImp struct {
	horarioRepositorio ports.HorarioRepositorio
}

//New is a function which creates a new [HorarioServiceImp]
func New(horarioRepositorio ports.HorarioRepositorio) *HorarioServiceImp {
	return &HorarioServiceImp{horarioRepositorio: horarioRepositorio}
}

//GetAvaiabledHours is a function which returns a set of [AvailableHours]
//given a completed [Terna] (not null fields)
func (srv *HorarioServiceImp) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {
	res, err := srv.horarioRepositorio.GetAvailableHours(terna)
	if err != nil {
		return []domain.AvailableHours{}, err
	}

	return res, nil
}

func (srv *HorarioServiceImp) CreateNewEntry(entry domain.Entry) (string, error) {
	err := entry.IsValid()
	if err != nil {
		return "", err
	}

	//Check if the entry has valid time interval
	if entry.Init.IsLaterThan(entry.End) {
		return "", apperrors.ErrInvalidInput
	}

	err = srv.horarioRepositorio.CreateNewEntry(entry)
	if err != nil {
		return "", apperrors.ErrInternal
	}
	return time.Now().Format("02/01/2006"), nil
}

func (srv *HorarioServiceImp) ListAllDegrees() ([]domain.DegreeDescription, error) {
	list, err := srv.horarioRepositorio.ListAllDegrees()
	return list, err
}

func (srv *HorarioServiceImp) UpdateScheduler(entries []domain.Entry, terna domain.Terna) (string, error) {
	var lastMod string
	err := srv.horarioRepositorio.DeleteAllEntries(terna)
	if err != nil {
		return "", apperrors.ErrSql
	}

	lastMod = time.Now().Format("02/01/2006")

	for i, e := range entries {
		//add
		date, err := srv.CreateNewEntry(e)
		if err != nil {
			return "", apperrors.ErrSql
		}
		if len(entries)-1 == i {
			lastMod = date
		}

	}
	return lastMod, nil
}

func (srv *HorarioServiceImp) GetEntries(terna domain.Terna) ([]domain.Entry, error) {

	if terna.Degree == "" || terna.Year == 0 || terna.Group == "" {
		return []domain.Entry{}, apperrors.ErrInvalidInput
	}
	entries, err := srv.horarioRepositorio.GetEntries(terna)
	if err != nil {
		return []domain.Entry{}, apperrors.ErrNotFound
	}
	return entries, nil

}

func (srv *HorarioServiceImp) GetICS(terna domain.Terna) (string, error) {
	if terna.Degree == "" || terna.Year == 0 || terna.Group == "" {
		return "", apperrors.ErrInvalidInput
	}
	entries, err := srv.horarioRepositorio.GetEntries(terna)
	if err != nil {
		return "", apperrors.ErrSql
	}
	cal := ics.NewCalendar()
	t := time.Now()
	month := t.Month()
	year := t.Year()
	if(month<8){
		month = 2
	} else {
		month = 9
	}
	taux := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for i, entry := range entries {
		event := cal.AddEvent(fmt.Sprintf("%d@unizar.es", i))
		event.SetSummary(entry.Subject.Name)
		day := (8-int(taux.Weekday()))%7 + entry.Weekday + 1
		t1 := time.Date(year, month, day, entry.Init.Hour, entry.Init.Min, 0, 0, t.Location())
		event.SetStartAt(t1)
		t2 := time.Date(year, month, day, entry.End.Hour, entry.End.Min, 0, 0, t.Location())
		event.SetEndAt(t2)
		event.AddRrule("FREQ=DAILY;INTERVAL=7")
		i++
	}
	return cal.Serialize(), nil
}

//3->Code subject, 4->Name subject, 11->Code Degree, 12->Name degree
//17->Year (1,2...), 23->nº groups, 30->Hours t1, 32->Hours t2, 34->Hours t3
//71->nº subgroups t1, 72->nº subgroups t1
func (srv *HorarioServiceImp) UpdateByCSV(csv string) (bool, error){
	lines := strings.Split(csv, "\n")
	prevDegree := 0
	for i, actLine := range lines {
		if(i<4){
			continue
		}
		cells := strings.Split(actLine, ";")
		subjectId, _ := strconv.Atoi(cells[3])
		subjectName := cells[4]
		degreeId, _ := strconv.Atoi(cells[11])
		degreeName := cells[12]
		year, _ := strconv.Atoi(cells[17])
		nGroups, _ := strconv.Atoi(cells[23])
		nGroupsT2, _ := strconv.Atoi(cells[71])
		nGroupsT2 /= 2
		nGroupsT3, _ := strconv.Atoi(cells[72])
		nGroupsT3 /= 2
		hoursT1, _ := strconv.Atoi(cells[30])
		hoursT2, _ := strconv.Atoi(cells[32])
		hoursT3, _ := strconv.Atoi(cells[34])
		if prevDegree != degreeId {
			srv.horarioRepositorio.CreateNewDegree(degreeId, degreeName)
			prevDegree = degreeId
		}
		srv.horarioRepositorio.CreateNewYear(year, degreeId)
		actYearId := degreeId * 10 + year
		ok, err := srv.horarioRepositorio.CreateNewSubject(subjectId, subjectName, degreeId)
		if !ok {
			return false, err
		}
		for j:=1;j<=nGroups;j++{
			srv.horarioRepositorio.CreateNewGroup(j, actYearId)
			ok, err = srv.horarioRepositorio.CreateNewHour(hoursT1*100,hoursT1*100,subjectId,actYearId*10+j,domain.THEORICAL,"","")
			if !ok {
				return false, err
			}
			for k:=1;k<=nGroupsT2;k++{
				ok, err = srv.horarioRepositorio.CreateNewHour(hoursT2*100,hoursT2*100,subjectId,actYearId*10+j,domain.EXERCISES,strconv.Itoa(k),"")
				if !ok {
					return false, err
				}
			}
			for k:=1;k<=nGroupsT3;k++{
				ok, err = srv.horarioRepositorio.CreateNewHour(hoursT3*100,hoursT3*100,subjectId,actYearId*10+j,domain.PRACTICES,strconv.Itoa(k),"a")
				if !ok {
					return false, err
				}
			}
		}
	}
	return true, nil
}