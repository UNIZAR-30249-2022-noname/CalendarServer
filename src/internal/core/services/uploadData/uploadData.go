package uploaddata

import (
	"fmt"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

//MonitoringServiceImp is the implemetation of [SchedulerService] interface.
type UploadDataServiceImp struct {
	uploadDataRepository ports.UploadDataRepository
}

//New is a function which creates a new [SchedulerServiceImp]
func New(uploadDataRepository ports.UploadDataRepository) *UploadDataServiceImp {
	return &UploadDataServiceImp{uploadDataRepository: uploadDataRepository}
}

func (srv *UploadDataServiceImp) UpdateByCSV(csv, privileges string) (bool, error) {
	if (privileges == "janitor"){
		fmt.Println("Janitor")
	return srv.uploadDataRepository.UpdateSpacesByCSV(csv)
	} else if (privileges == "coordinator") {
		fmt.Println("Coordinator")
		return srv.uploadDataRepository.UpdateDegreesByCSV(csv)
	 } else {
		fmt.Println("Other:" + privileges)
		 return false, apperrors.ErrNoRowsAffected
		}
}