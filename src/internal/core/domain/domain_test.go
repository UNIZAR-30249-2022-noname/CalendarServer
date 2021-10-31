package domain_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
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
		result := tt.args.h1.IsLaterThan(tt.args.h2)
		assert.Equal(t, tt.want, result)
	}

}
