
# 📦 Documentação de Testes

Este documento descreve os testes unitários implementados para o pacote `service`, incluindo estrutura do projeto, bibliotecas utilizadas, explicação dos testes existentes e instruções para escrever novos testes.

---

## 🗂 Estrutura do Projeto

```
project-root/
│
├── service/
│   └── statisticsservice_test.go  # Arquivo de testes unitários
│
├── pkg/
│   └── database/
│       └── connection.go  # Define a função ConnectDB()
```

⚠️ **Importante**: Certifique-se de que o banco esteja ativo e populado com os dados exigidos pelos testes (usuários, projetos, papéis, etc.).

---

## 📚 Bibliotecas Utilizadas

| Biblioteca                         | Uso                                                                 |
|------------------------------------|---------------------------------------------------------------------|
| `testing` (Go padrão)              | Estrutura base para definição de testes (`t *testing.T`)            |
| `github.com/stretchr/testify/assert` | Biblioteca externa para asserções mais legíveis e robustas         |

Para instalar o `testify`:
```bash
go get github.com/stretchr/testify
```

---

## ✅ Testes Implementados

### 🔹 `TestGetListCardTags_AdminUser`
- Testa a função `GetListCardTags` para um usuário administrador (`idUser = 0`).
- Verifica:
  - Se o resultado é retornado corretamente com dados válidos.
  - Se a função trata erros adequadamente quando recebe IDs inválidos.

---

### 🔹 `TestGetListCardTags_OperatorUser`
- Testa `GetListCardTags` para um operador (`idUser > 0`).
- Verifica comportamento com dados válidos e inválidos.

---

### 🔹 `TestGetMetricsRole_Admin`
- Testa a função `GetMetricsRole` para admin.
- Garante que:
  - Retorna `http.StatusOK` e resposta válida quando os dados estão corretos.
  - Retorna `http.StatusBadRequest` com erro no caso de IDs inexistentes.

---

### 🔹 `TestGetMetrics_AdminRole`
- Testa `GetMetrics` com role “ADMIN”.
- Garante retorno correto com `idRole` real do banco.
- Também cobre caso com dados inválidos.

---

### 🔹 `TestGetMetrics_OperatorRole`
- Testa `GetMetrics` com role “OPERADOR” (ou outro papel não-admin).
- Verifica resposta correta e tratamento de erro com parâmetros inválidos.

---

## 🛠 Como Criar um Teste Unitário em Go

### Exemplo básico com `testing` e `assert`
```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMinhaFuncao(t *testing.T) {
    // Chamada da função que será testada
    resultado, err := MinhaFuncao("entrada")

    // Verificações com assert
    assert.Nil(t, err, "Esperava erro nulo")
    assert.Equal(t, "resultado esperado", resultado)
}
```

### Estrutura de um teste padrão
1. Nome da função de teste: `Test<NomeFuncionalidade>`
2. Inicializa os dados de entrada.
3. Chama a função alvo.
4. Usa `assert.Nil`, `assert.NotNil`, `assert.Equal`, etc., para validar.
5. Opcional: usa `t.Logf(...)` para debug/log.

---

## 🔧 Pré-Requisitos

- Go 1.18 ou superior  
- Banco de dados configurado e populado com os dados esperados  
- Módulo Go inicializado (`go mod init`)  
- Dependências instaladas:  
```bash
go mod tidy
```

---

## 🚀 Como Executar os Testes

Para rodar todos os testes do projeto:
```bash
go test ./...
```

Para rodar apenas os testes do arquivo `service_test.go`:
```bash
go test ./service -v
```

---

## 🧪 Dicas para Expandir os Testes

- Sempre valide casos **felizes (happy path)** e **casos de erro**.
- Use `t.Run("descrição", func(t *testing.T) {...})` para subtestes.
- Crie funções auxiliares para montar dados repetitivos (ex: datas, IDs).
- Prefira `assert` ao invés de `if t.Errorf` para clareza.

---

## 📌 Conclusão

Este documento serve como base para quem desejar manter ou expandir os testes do serviço. Siga o padrão estabelecido e mantenha sempre cobertura para os principais fluxos e erros possíveis.
