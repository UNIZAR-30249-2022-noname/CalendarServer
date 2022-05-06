package users_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/users"
	mock_ports "github.com/D-D-EINA-Calendar/CalendarServer/src/mocks/mockups"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	usersRepository *mock_ports.MockUsersRepository
}

func TestGetCredentials(t *testing.T) {

	username := "admin"
	credentials := domain.User{Name: username, Privileges: "professor"}

	// · Test · //
	type args struct {
		username string
	}
	type want struct {
		result domain.User
		err    error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{

		{
			name: "Should return credentials",
			args: args{username: username},
			want: want{result: credentials, err: nil},
			mocks: func(m mocks) {
				m.usersRepository.EXPECT().GetCredentials(username).Return(credentials, nil)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Prepare

			m := mocks{
				usersRepository: mock_ports.NewMockUsersRepository(gomock.NewController(t)),
			}

			tt.mocks(m)

			service := users.New(m.usersRepository)

			//Execute
			result, err := service.GetCredentials(tt.args.username)

			//Verify
			if tt.want.err != nil && err != nil {
				assert.Equal(t, tt.want.err.Error(), err.Error())
			}

			assert.Equal(t, tt.want.result, result)

		})

	}

}
