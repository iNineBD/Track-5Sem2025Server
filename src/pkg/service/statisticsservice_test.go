package service

import (
	"inine-track/pkg/database"
	"log"
	"testing"
)

func TestGetMetrics(t *testing.T) {
	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	type args struct {
		IDProject int
		Data1     string
		Data2     string
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{
			name: "VALID",
			args: args{
				1,
				"2025-04-01",
				"2025-04-30",
			},
			wantStatus: 200,
		}, {
			name: "INVALID",
			args: args{
				999999999,
				"2025-04-01",
				"2025-04-30",
			},
			wantStatus: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotStatus, _ := GetMetrics(tt.args.IDProject, tt.args.Data1, tt.args.Data2)

			// Verifica o status retornado
			if gotStatus != tt.wantStatus {
				t.Errorf("GetCardsPerStatus() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}

		})
	}
}
