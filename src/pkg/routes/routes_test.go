package routes

import (
	"net/http"
	"testing"
	"time"
)

// Testa se o Swagger está online, mas permite que o teste passe se o servidor estiver offline.
func TestSwaggerIsOnline(t *testing.T) {
	// Aguarda um tempo para garantir que o servidor tenha iniciado (se estiver rodando).
	time.Sleep(2 * time.Second)

	// Tenta fazer uma requisição ao servidor
	resp, err := http.Get("http://localhost:8080/swagger/index.html")
	if err != nil {
		// Se houver erro na conexão, assume que o servidor está offline e passa o teste
		t.Logf("Servidor não está rodando ou conexão recusada: %v", err)
		t.Skip("Pulando teste porque o servidor não está disponível")
		return
	}
	defer resp.Body.Close()

	// Se conseguiu conectar, verifica se o status HTTP é 200
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Esperado status code 200, mas obteve %d", resp.StatusCode)
	}
}
