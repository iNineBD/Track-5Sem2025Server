package service

import (
	"inine-track/pkg/database"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	database.ConnectDB()
}

// Teste para GetListCardTags com admin (idUser = 0)
func TestGetListCardTags_AdminUser(t *testing.T) {
	IDProject := int64(1648306)
	idUser := int64(0)
	data1 := time.Now().AddDate(0, 0, -100)
	data2 := time.Now()

	result, err := GetListCardTags(IDProject, data1, data2, idUser)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	t.Logf("Resultado (Admin): %+v", result)

	IDProjectComErro := int64(0)
	idUserComErro := int64(9999999)
	data1ComErro := time.Now().AddDate(0, 0, -100)
	data2ComErro := time.Now()

	result, err = GetListCardTags(IDProjectComErro, data1ComErro, data2ComErro, idUserComErro)

	assert.Nil(t, err)
	assert.Empty(t, result)
	t.Logf("Resultado (Admin): %+v", result)
}

// Teste para GetListCardTags com operador (idUser > 0)
func TestGetListCardTags_OperatorUser(t *testing.T) {
	IDProject := int64(1648306)
	idUser := int64(1648306)
	data1 := time.Now().AddDate(0, 0, -100)
	data2 := time.Now()

	result, err := GetListCardTags(IDProject, data1, data2, idUser)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	t.Logf("Resultado (Operador): %+v", result)

	IDProjectComErro := int64(0)
	idUserComErro := int64(9999999)
	data1ComErro := time.Now().AddDate(0, 0, -100)
	data2ComErro := time.Now()

	result, err = GetListCardTags(IDProjectComErro, data1ComErro, data2ComErro, idUserComErro)

	assert.Nil(t, err)
	assert.Empty(t, result)
	t.Logf("Resultado (Operador): %+v", result)
}

// Teste para GetMetricsRole com admin
func TestGetMetricsRole_Admin(t *testing.T) {
	IDProject := int64(1648306)
	idUser := int64(0)
	data1 := time.Now().AddDate(0, 0, -100)
	data2 := time.Now()

	status, response := GetMetricsRole(IDProject, data1, data2, idUser)

	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, response["success"])
	t.Logf("Response (Admin): %+v", response)

	IDProjectComErro := int64(9999999)
	idUserComErro := int64(9999999)
	data1ComErro := time.Now().AddDate(0, 0, -100)
	data2ComErro := time.Now()

	status, response = GetMetricsRole(IDProjectComErro, data1ComErro, data2ComErro, idUserComErro)

	assert.NotNil(t, response["error"])
	assert.Equal(t, http.StatusBadRequest, status)
	t.Logf("Resultado (Admin): %+v", response)

}

// Teste para GetMetrics com ADMIN
func TestGetMetrics_AdminRole(t *testing.T) {
	IDProject := int64(1648306)
	data1 := time.Now().AddDate(-1, -1, -100).Format("2006-01-02")
	data2 := time.Now().Format("2006-01-02")
	idUser := int64(2)
	idRole := int64(9965612) // deve estar no banco com NameRole = "ADMIN"

	status, response := GetMetrics(IDProject, data1, data2, idUser, idRole)

	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, response["success"])
	t.Logf("Response (Admin Role): %+v", response)

	IDProjectComErro := int64(9999999)
	idUserComErro := int64(9999999)
	data1ComErro := time.Now().AddDate(0, 0, -100)
	data2ComErro := time.Now()

	status, response = GetMetricsRole(IDProjectComErro, data1ComErro, data2ComErro, idUserComErro)

	assert.NotNil(t, response["error"])
	assert.Equal(t, http.StatusBadRequest, status)
	t.Logf("Resultado (Admin Role): %+v", response)
}

// Teste para GetMetrics com OPERADOR
func TestGetMetrics_OperatorRole(t *testing.T) {
	IDProject := int64(1648306)
	data1 := time.Now().AddDate(-1, -1, -100).Format("2006-01-02")
	data2 := time.Now().Format("2006-01-02")
	idUser := int64(758625)
	idRole := int64(10037040) // deve estar no banco com outro nome que não seja ADMIN ou GESTOR

	status, response := GetMetrics(IDProject, data1, data2, idUser, idRole)

	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, response["success"])
	t.Logf("Response (Operador Role): %+v", response)

	IDProjectComErro := int64(9999999)
	idUserComErro := int64(9999999)
	data1ComErro := time.Now().AddDate(0, 0, -100)
	data2ComErro := time.Now()

	status, response = GetMetricsRole(IDProjectComErro, data1ComErro, data2ComErro, idUserComErro)

	assert.NotNil(t, response["error"])
	assert.Equal(t, http.StatusBadRequest, status)
	t.Logf("Resultado (Operador Role): %+v", response)
}
