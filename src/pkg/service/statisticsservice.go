package service

import (
	"errors"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service/utils"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetMetrics(IDProject int64, data1 string, data2 string) (status int, response gin.H) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err}
	}

	listCardsPerTag, err2 := GetListCardTags(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsPerUser, err2 := GetListCardsPerUser(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsPerStatus, err2 := GetListCardsPerStatus(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsRework, err2 := GetListCardsRework(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsStarted, err2 := GetListCardsStarted(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsFinished, err2 := GetListCardsFinished(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsTimeExecution, err2 := GetListCardsTimeExecution(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	response = gin.H{"success": statisticsdto.GetStatisticsResponse{TagData: listCardsPerTag, UserData: listCardsPerUser,
		StatusData: listCardsPerStatus, ReworkCards: listCardsRework, StartedCards: listCardsStarted, FinishedCards: listCardsFinished,
		ExecutionCards: listCardsTimeExecution}}

	return http.StatusOK, response
}

func GetListCardTags(IDProject int64, data1 time.Time, data2 time.Time) (listCardsPerTag []statisticsdto.TagData, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_por_tag($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerTag)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards por tag"}
	}

	return listCardsPerTag, nil
}

func GetListCardsPerUser(IDProject int64, data1 time.Time, data2 time.Time) (listCardsPerUser []statisticsdto.UserData, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_por_colaborador($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerUser)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar as cards por colaborador"}
	}

	return listCardsPerUser, nil
}

func GetListCardsPerStatus(IDProject int64, data1 time.Time, data2 time.Time) (listCardsPerStatus []statisticsdto.StatusData, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_por_status($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerStatus)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar as cards por status"}
	}

	return listCardsPerStatus, nil
}

func GetListCardsRework(IDProject int64, data1 time.Time, data2 time.Time) (listCardsRework []statisticsdto.ReworkCards, err gin.H) {

	result := database.DB.Raw(`select * from get_retrabalhos($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsRework)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de retrabalho por card"}
	}

	return listCardsRework, nil
}

func GetListCardsStarted(IDProject int64, data1 time.Time, data2 time.Time) (listCardsStarted []statisticsdto.StartedCards, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsStarted)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards criados por projeto"}
	}

	return listCardsStarted, nil
}

func GetListCardsFinished(IDProject int64, data1 time.Time, data2 time.Time) (listCardsFinished []statisticsdto.FinishedCards, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsFinished)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards finalizados por projet"}
	}

	return listCardsFinished, nil
}

func GetListCardsTimeExecution(IDProject int64, data1 time.Time, data2 time.Time) (listCardsTimeExecution []statisticsdto.TimeExecutionCards, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsTimeExecution)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar o tempo de execução dos cards"}
	}

	return listCardsTimeExecution, nil
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
				1, time.Now().AddDate(0, 0, -7), time.Now(),
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
				1, time.Now().AddDate(0, 0, -30), time.Now(),
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
				2, time.Now().AddDate(0, -1, 0), time.Now(),
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
