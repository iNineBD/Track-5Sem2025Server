package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/service/utils"
	"log"
	"net/http"
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
