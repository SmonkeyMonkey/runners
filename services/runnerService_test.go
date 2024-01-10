package services

import (
	"net/http"
	"testing"

	"github.com/smonkeymonkey/marathon_runners/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateRunner(t *testing.T) {
	tests := []struct {
		name   string
		runner *models.Runner
		want   *models.ResponseError
	}{
		{
			name: "Invalid_First_Name",
			runner: &models.Runner{
				LastName: "Smith",
				Age:      35,
				Country:  "US",
			},
			want: &models.ResponseError{
				Message: "Invalid first name",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_Last_name",
			runner: &models.Runner{
				FirstName: "John",
				Age:       31,
				Country:   "US",
			},
			want: &models.ResponseError{
				Message: "Invalid last name",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_Age",
			runner: &models.Runner{
				FirstName: "John",
				LastName:  "Smith",
				Age:       355,
				Country:   "US",
			},
			want: &models.ResponseError{
				Message: "Invalid age",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_Country",
			runner: &models.Runner{
				FirstName: "John",
				LastName:  "Smith",
				Age:       35,
			},
			want: &models.ResponseError{
				Message: "Invalid country",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Valid runner",
			runner: &models.Runner{
				FirstName: "John",
				LastName:  "Smith",
				Age:       35,
				Country:   "US",
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responseErr := validateRunner(test.runner)
			assert.Equal(t, test.want, responseErr)
		})
	}
}
