package service

import (
	"inine-track/pkg/database"
	"log"
	"testing"
)

func TestGetCardsPerStatus(t *testing.T) {
	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	type args struct {
		IDProject int
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{
			name: "VALID",
			args: args{
				4,
			},
			wantStatus: 200,
		}, {
			name: "INVALID",
			args: args{
				999999999,
			},
			wantStatus: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotStatus, _ := GetCardsPerStatus(tt.args.IDProject)

			// Verifica o status retornado
			if gotStatus != tt.wantStatus {
				t.Errorf("GetCardsPerStatus() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}

		})
	}
}

func TestGetCardsPerUser(t *testing.T) {

	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	type args struct {
		IDProject int
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		// TODO: Add test cases.
		{
			name: "VALID",
			args: args{
				4,
			},
			wantStatus: 200,
		}, {
			name: "INVALID",
			args: args{
				999999999,
			},
			wantStatus: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, _ := GetCardsPerUser(tt.args.IDProject)
			if gotStatus != tt.wantStatus {
				t.Errorf("GetCardsPerUser() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestGetCardsPerTag(t *testing.T) {

	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}
	type args struct {
		IDProject int
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		// TODO: Add test cases.
		{
			name: "VALID",
			args: args{
				4,
			},
			wantStatus: 200,
		}, {
			name: "INVALID",
			args: args{
				999999999,
			},
			wantStatus: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, _ := GetCardsPerTag(tt.args.IDProject)
			if gotStatus != tt.wantStatus {
				t.Errorf("GetCardsPerTag() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
