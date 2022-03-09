package users

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

type UsersServiceImp struct {
	usersRepository ports.UsersRepository
}

func New(usersRepository ports.UsersRepository) *UsersServiceImp {
	return &UsersServiceImp{usersRepository: usersRepository}
}

func (svc *UsersServiceImp) GetCredentials(username string) (domain.User, error) {
	return svc.usersRepository.GetCredentials(username)
}
