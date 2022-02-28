package domain_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestIsLaterThan(t *testing.T) {

	type args struct {
		h1 domain.Hour
		h2 domain.Hour
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "One hour  earlier",
			args: args{h1: domain.NewHour(1, 0), h2: domain.NewHour(2, 0)},
			want: false,
		},
		{
			name: "One hour  later",
			args: args{h1: domain.NewHour(3, 0), h2: domain.NewHour(2, 0)},
			want: true,
		},

		{
			name: "One minute earlier",
			args: args{h1: domain.NewHour(1, 0), h2: domain.NewHour(1, 1)},
			want: false,
		},
		{
			name: "One minute later",
			args: args{h1: domain.NewHour(1, 2), h2: domain.NewHour(1, 1)},
			want: true,
		},
		{
			name: "One minute earlier diferent hour",
			args: args{h1: domain.NewHour(1, 59), h2: domain.NewHour(2, 0)},
			want: false,
		},
		{
			name: "One minute later diferent hour",
			args: args{h1: domain.NewHour(2, 0), h2: domain.NewHour(1, 59)},
			want: true,
		},
		{
			name: "Same hour",
			args: args{h1: domain.NewHour(2, 0), h2: domain.NewHour(2, 0)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.h1.IsLaterThan(tt.args.h2)
			assert.Equal(t, tt.want, result)

		})

	}

}

func TestSubjectValid(t *testing.T) {
	tests := []struct {
		name string
		args domain.Subject
		want error
	}{
		{
			name: "valid",
			args: domain.Subject{Name: "a", Kind: constants.THEORICAL},
			want: nil,
		},
		{
			name: "no empty name",
			args: domain.Subject{Kind: constants.THEORICAL},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "no empty name",
			args: domain.Subject{Name: "a"},
			want: apperrors.ErrInvalidInput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.IsValid()
			assert.Equal(t, tt.want, result)
		})

	}

}

func TestEntryValid(t *testing.T) {

	tests := []struct {
		name string
		args domain.Entry
		want error
	}{
		{
			name: "init Hour lacks",
			args: domain.Entry{
				End: domain.NewHour(2, 0),
				Subject: domain.Subject{
					Kind: constants.THEORICAL,
					Name: "Prog 1",
				},
				Room: domain.Room{Name: "1"},
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "End Hour lacks",
			args: domain.Entry{
				Init: domain.NewHour(2, 0),
				Subject: domain.Subject{
					Kind: constants.THEORICAL,
					Name: "Prog 1",
				},
				Room: domain.Room{Name: "1"},
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "subject  lacks",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Room: domain.Room{Name: "1"},
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "Subject  empty name",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Subject: domain.Subject{
					Kind: constants.THEORICAL,
				},
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "Subject  empty kind",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Subject: domain.Subject{
					Name: "Prog 1",
				},
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "Subject valid",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Subject: domain.Subject{
					Kind: constants.THEORICAL,
					Name: "Prog 1",
				},
			},
			want: nil,
		},
		{
			name: "Practices lacks week ",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Room: domain.Room{Name: "1"},
				Subject: domain.Subject{
					Kind: constants.PRACTICES,
					Name: "Prog 1",
				},
				Group: "1",
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "Practices lacks group ",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Room: domain.Room{Name: "1"},
				Subject: domain.Subject{
					Kind: constants.PRACTICES,
					Name: "Prog 1",
				},
				Week: "a",
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "Practices valid ",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Room: domain.Room{Name: "1"},
				Subject: domain.Subject{
					Kind: constants.PRACTICES,
					Name: "Prog 1",
				},
				Week:  "a",
				Group: "1",
			},
			want: nil,
		},
		{
			name: "Exercises lacks group ",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Room: domain.Room{Name: "1"},
				Subject: domain.Subject{
					Kind: constants.EXERCISES,
					Name: "Prog 1",
				},
			},
			want: apperrors.ErrInvalidInput,
		},
		{
			name: "Exercises valid ",
			args: domain.Entry{
				Init: domain.NewHour(1, 0),
				End:  domain.NewHour(2, 0),
				Room: domain.Room{Name: "1"},
				Subject: domain.Subject{
					Kind: constants.EXERCISES,
					Name: "Prog 1",
				},

				Group: "1",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.IsValid()
			assert.Equal(t, tt.want, result)

		})

	}

}
