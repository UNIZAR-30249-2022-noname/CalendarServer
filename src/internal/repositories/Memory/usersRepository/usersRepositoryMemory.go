package usersrepositorymemory

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type UsersRepositoryMemory struct {
}

func New() *UsersRepositoryMemory {
	return &UsersRepositoryMemory{}
}
func (repo *UsersRepositoryMemory) GetCredentials(username string) (domain.User, error) {
	var privileges string
	if username == "785370" {
		privileges = "professor"
	} else {
		privileges = "none"
	}
	return domain.User{Name: username, Privileges: privileges}, nil
}
