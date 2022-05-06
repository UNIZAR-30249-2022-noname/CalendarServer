package uploaddata

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

//MonitoringServiceImp is the implemetation of [SchedulerService] interface.
type UploadDataServiceImp struct {
	uploadDataRepository ports.UploadDataRepository
}

//New is a function which creates a new [SchedulerServiceImp]
func New(uploadDataRepository ports.UploadDataRepository) *UploadDataServiceImp {
	return &UploadDataServiceImp{uploadDataRepository: uploadDataRepository}
}

func (srv *UploadDataServiceImp) UpdateByCSV(csv string) (bool, error) {
	return srv.uploadDataRepository.UpdateByCSV(csv)
}