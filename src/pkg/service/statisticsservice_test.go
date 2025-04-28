package service

import (
	"errors"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"

	//	"inine-track/pkg/service" -- Ciclo desnecessário de importação
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
				1648306,
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
				IDProject: 1648306,
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
				IDProject: 1648306,
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
				IDProject: 1648306,
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

func TestGetListCardsRework(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	tests := []struct {
		name string
		args struct {
			IDProject    int64
			data1, data2 time.Time
		}
		wantReworkCards []statisticsdto.ReworkCards
		wantErr         gin.H
	}{
		{
			name: "Erro no banco",
			args: struct {
				IDProject    int64
				data1, data2 time.Time
			}{
				1648306, time.Now().AddDate(0, 0, -7), time.Now(),
			},
			wantReworkCards: nil,
			wantErr:         gin.H{"error": "erro ao retornar a quantidade de retrabalho por card"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("select \\* from get_retrabalhos").
				WillReturnError(errors.New("erro no banco simulado"))

			got, err := GetListCardsRework(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(got, tt.wantReworkCards) {
				t.Errorf("got = %v, want %v", got, tt.wantReworkCards)
			}
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("gotErr = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetListCardsStarted(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	tests := []struct {
		name string
		args struct {
			IDProject    int64
			data1, data2 time.Time
		}
		wantStartedCards []statisticsdto.StartedCards
		wantErr          gin.H
	}{
		{
			name: "Erro no banco",
			args: struct {
				IDProject    int64
				data1, data2 time.Time
			}{
				1648306, time.Now().AddDate(0, -1648306, 0), time.Now(),
			},
			wantStartedCards: nil,
			wantErr:          gin.H{"error": "erro ao retornar a quantidade de cards criados por projeto"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("select \\* from get_qtd_cards_criados_por_projeto").
				WillReturnError(errors.New("simulação de erro"))

			got, err := GetListCardsStarted(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(got, tt.wantStartedCards) {
				t.Errorf("got = %v, want %v", got, tt.wantStartedCards)
			}
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("gotErr = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetListCardsFinished(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	tests := []struct {
		name string
		args struct {
			IDProject    int64
			data1, data2 time.Time
		}
		wantFinishedCards []statisticsdto.FinishedCards
		wantErr           gin.H
	}{
		{
			name: "Erro no banco",
			args: struct {
				IDProject    int64
				data1, data2 time.Time
			}{
				1648306, time.Now().AddDate(0, 0, -30), time.Now(),
			},
			wantFinishedCards: nil,
			wantErr:           gin.H{"error": "erro ao retornar a quantidade de cards finalizados por projet"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("select \\* from get_qtd_cards_criados_por_projeto").
				WillReturnError(errors.New("simulação de erro"))

			got, err := GetListCardsFinished(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(got, tt.wantFinishedCards) {
				t.Errorf("got = %v, want %v", got, tt.wantFinishedCards)
			}
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("gotErr = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetListCardsTimeExecution(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer db.Close()

	database.DB, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	tests := []struct {
		name string
		args struct {
			IDProject    int64
			data1, data2 time.Time
		}
		wantTimeExecution []statisticsdto.TimeExecutionCards
		wantErr           gin.H
	}{
		{
			name: "Erro no banco",
			args: struct {
				IDProject    int64
				data1, data2 time.Time
			}{
				2, time.Now().AddDate(0, -1648306, 0), time.Now(),
			},
			wantTimeExecution: nil,
			wantErr:           gin.H{"error": "erro ao retornar o tempo de execução dos cards"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("select \\* from get_qtd_cards_criados_por_projeto").
				WillReturnError(errors.New("falha intencional"))

			got, err := GetListCardsTimeExecution(tt.args.IDProject, tt.args.data1, tt.args.data2)

			if !reflect.DeepEqual(got, tt.wantTimeExecution) {
				t.Errorf("got = %v, want %v", got, tt.wantTimeExecution)
			}
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("gotErr = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
