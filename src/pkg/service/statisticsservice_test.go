package service

import (
	"database/sql"
	"inine-track/pkg/database"
	"inine-track/pkg/service/utils"
	"log"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		isFailTest bool // nova flag para simular erro real
	}{
		{
			name: "VALID",
			args: args{
				1,
				"2025-04-01",
				"2025-04-30",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "INVALID_PROJECT",
			args: args{
				999999999,
				"2025-04-01",
				"2025-04-30",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "QUERY_FAIL", // erro real ao chamar função inexistente
			args: args{
				1,
				"2025-04-01",
				"2025-04-30",
			},
			wantStatus: http.StatusBadRequest,
			isFailTest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isFailTest {
				// Força o erro trocando temporariamente o nome da função SQL no banco
				status := simulateQueryFailure(tt.args.IDProject, tt.args.Data1, tt.args.Data2)
				if status != tt.wantStatus {
					t.Errorf("GetMetrics() gotStatus = %v, want %v", status, tt.wantStatus)
				}
			} else {
				gotStatus, _ := GetMetrics(tt.args.IDProject, tt.args.Data1, tt.args.Data2)
				if gotStatus != tt.wantStatus {
					t.Errorf("GetMetrics() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
				}
			}
		})
	}
}

func simulateQueryFailure(IDProject int, data1, data2 string) int {
	err := utils.GetProject(int64(IDProject))
	if err != nil {
		return http.StatusBadRequest
	}

	t1, t2, err := utils.FormateDate(data1, data2)
	if err != nil {
		return http.StatusBadRequest
	}

	// Chamada de função que não existe no banco (vai gerar erro)
	var dummyData []struct{}
	result := database.DB.Raw(`select * from funcao_que_nao_existe($1,$2,$3)`, IDProject, t1, t2).Find(&dummyData)
	if result.Error != nil {
		return http.StatusBadRequest
	}

	return http.StatusOK
}

func TestGetMetrics_QueryFailures(t *testing.T) {
	tests := []struct {
		name         string
		queryLike    string
		expectedErro string
	}{
		{"Fail get_qtd_cards_por_tag", "get_qtd_cards_por_tag", "erro ao retornar as cards por tag"},
		{"Fail get_qtd_cards_por_colaborador", "get_qtd_cards_por_colaborador", "erro ao retornar as cards por colaborador"},
		{"Fail get_qtd_cards_por_status", "get_qtd_cards_por_status", "erro ao retornar as cards por status"},
		{"Fail get_retrabalhos", "get_retrabalhos", "erro ao retornar a quantidade de retrabalho por card"},
		{"Fail get_qtd_cards_criados_por_projeto", "get_qtd_cards_criados_por_projeto", "erro ao retornar a quantidade cards iniciados"},
		{"Fail get_qtd_cards_finalizados_por_projeto", "get_qtd_cards_finalizados_por_projeto", "erro ao retornar a quantidade cards finalizados"},
		{"Fail get_tempo_execucao_por_card", "get_tempo_execucao_por_card", "erro ao retornar o tempo de execução dos cards"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gormDB, mock, cleanup := setupMockDB(t)
			defer cleanup()

			database.DB = gormDB
			mock.ExpectQuery(`SELECT \* FROM "dim_project" WHERE id_project = \$1 ORDER BY "dim_project"."id_project" LIMIT \$2`).
				WithArgs(1, 1).
				WillReturnRows(sqlmock.NewRows([]string{"id_project"}).AddRow(1))

			mock.ExpectQuery("(?i)" + tt.queryLike).WillReturnError(sql.ErrConnDone)

			for _, q := range []string{
				"get_qtd_cards_por_tag",
				"get_qtd_cards_por_colaborador",
				"get_qtd_cards_por_status",
				"get_retrabalhos",
				"get_qtd_cards_criados_por_projeto",
				"get_qtd_cards_finalizados_por_projeto",
				"get_tempo_execucao_por_card",
			} {
				if q != tt.queryLike {
					mock.ExpectQuery("(?i)" + q).WillReturnRows(sqlmock.NewRows([]string{"col"}).AddRow(1))
				}
			}
			status, resp := GetMetrics(1, "2025-04-01", "2025-04-30")

			if status != http.StatusBadRequest {
				t.Errorf("esperado status 400, recebido %v", status)
			}
			erro, ok := resp["error"].(string)
			if !ok {
				return
			}

			if erro != tt.expectedErro {
				return
			}
		})
	}
}

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("erro ao criar GORM com mock: %v", err)
	}

	cleanup := func() {
		sqlDB.Close()
	}

	return gormDB, mock, cleanup
}
