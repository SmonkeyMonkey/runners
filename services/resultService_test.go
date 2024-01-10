package services

import (
	"net/http"
	"testing"

	"github.com/smonkeymonkey/marathon_runners/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateResult(t *testing.T) {
	tests := []struct {
		name   string
		result *models.Result
		want   *models.ResponseError
	}{
		{
			name: "Invalid_Runner_ID",
			result: &models.Result{
				ID:         "1",
				RaceResult: "02:01:23",
				Location:   "US",
				Year:       2024,
			},
			want: &models.ResponseError{
				Message: "Invalid runner ID",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_Race_Resultt",
			result: &models.Result{
				ID:       "1",
				RunnerID: "2",
				Location: "US",
				Year:     2024,
			},
			want: &models.ResponseError{
				Message: "Invalid race result",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_Locaiton",
			result: &models.Result{
				ID:         "1",
				RunnerID:   "2",
				RaceResult: "02:01:23",
				Year:       2024,
			},
			want: &models.ResponseError{
				Message: "Invalid location",
				Status:  http.StatusBadRequest,
			},
		},
		{
			name: "Invalid_Position",
			result: &models.Result{
				ID:         "1",
				RunnerID:   "2",
				RaceResult: "02:01:23",
				Location:   "US",
				Position:   -1,
				Year:       2024,
			},
			want: &models.ResponseError{
				Message: "Invalid position",
				Status:  http.StatusBadRequest,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responseErr := validateResult(test.result)
			assert.Equal(t, test.want, responseErr)
		})
	}
}
