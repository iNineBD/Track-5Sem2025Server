package service

import (
	"errors"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetMetrics(t *testing.T) {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	type args struct {
		IDProject int64
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
		},
		{
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
			gotStatus, gotResponse := GetMetrics(tt.args.IDProject, tt.args.Data1, tt.args.Data2)
			if gotStatus != tt.wantStatus {
				t.Errorf("GetMetrics() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotStatus == 200 {
				t.Logf("GetMetrics() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestGetListCardTags(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	type args struct {
		IDProject int64
		data1     time.Time
		data2     time.Time
	}
	tests := []struct {
		name                string
		args                args
		wantListCardsPerTag []statisticsdto.TagData
		wantErr             gin.H
	}{
		{
			name: "Erro ao executar query",
			args: args{
				IDProject: 1,
				data1:     time.Now().AddDate(0, 0, -30),
				data2:     time.Now(),
			},
			wantListCardsPerTag: nil,
			wantErr:             gin.H{"error": "erro ao retornar a quantidade de cards por tag"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("select \\* from get_qtd_cards_por_tag").
				WillReturnError(errors.New("erro de banco simulado"))

			gotListCardsPerTag, gotErr := GetListCardTags(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(gotListCardsPerTag, tt.wantListCardsPerTag) {
				t.Errorf("GetListCardTags() gotListCardsPerTag = %v, want %v", gotListCardsPerTag, tt.wantListCardsPerTag)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetListCardTags() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestGetListCardsPerUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	type args struct {
		IDProject int64
		data1     time.Time
		data2     time.Time
	}
	tests := []struct {
		name                 string
		args                 args
		wantListCardsPerUser []statisticsdto.UserData
		wantErr              gin.H
	}{
		{
			name: "Erro ao executar query",
			args: args{
				IDProject: 1,
				data1:     time.Now().AddDate(0, 0, -30),
				data2:     time.Now(),
			},
			wantListCardsPerUser: nil,
			wantErr:              gin.H{"error": "erro ao retornar as cards por colaborador"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectQuery("select \\* from get_qtd_cards_por_colaborador").
				WillReturnError(errors.New("erro de banco simulado"))

			gotListCardsPerUser, gotErr := GetListCardsPerUser(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(gotListCardsPerUser, tt.wantListCardsPerUser) {
				t.Errorf("GetListCardsPerUser() gotListCardsPerUser = %v, want %v", gotListCardsPerUser, tt.wantListCardsPerUser)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetListCardsPerUser() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestGetListCardsPerStatus(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	type args struct {
		IDProject int64
		data1     time.Time
		data2     time.Time
	}
	tests := []struct {
		name                   string
		args                   args
		wantListCardsPerStatus []statisticsdto.StatusData
		wantErr                gin.H
	}{
		{
			name: "Erro ao executar query",
			args: args{
				IDProject: 1,
				data1:     time.Now().AddDate(0, 0, -30),
				data2:     time.Now(),
			},
			wantListCardsPerStatus: nil,
			wantErr:                gin.H{"error": "erro ao retornar as cards por status"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectQuery("select \\* from get_qtd_cards_por_status").
				WillReturnError(errors.New("erro de banco simulado"))

			gotListCardsPerStatus, gotErr := GetListCardsPerStatus(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(gotListCardsPerStatus, tt.wantListCardsPerStatus) {
				t.Errorf("GetListCardsPerStatus() gotListCardsPerStatus = %v, want %v", gotListCardsPerStatus, tt.wantListCardsPerStatus)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetListCardsPerStatus() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
