
# 🧪 Documentação de Testes - Pacote `service`

Este documento descreve a estrutura, dependências e instruções para implementação e execução de testes no pacote `service` do projeto **inine-track**.

## 📂 Estrutura do Projeto

```
project-root/
│
├── service/
│   └── service_test.go       # Arquivo com testes unitários
│
├── pkg/
│   └── database/
│       └── connection.go     # Função ConnectDB() usada para conectar ao banco
```

⚠️ **Importante:** Certifique-se de que o banco esteja ativo e populado com os dados exigidos pelos testes (usuários, projetos, papéis, etc.).

## ✅ Testes Implementados

### `TestGetListCardTags_AdminUser`

Testa a função `GetListCardTags` para usuário admin (`idUser = 0`).  
Verifica se o retorno é válido e se o erro é tratado corretamente com IDs inválidos.

### `TestGetListCardTags_OperatorUser`

Testa a função `GetListCardTags` para usuário operador (`idUser > 0`).  
Também cobre cenários com entradas inválidas.

### `TestGetMetricsRole_Admin`

Testa `GetMetricsRole` para admin.  
Valida resposta esperada e tratamento de erro quando IDs não existem.

### `TestGetMetrics_AdminRole`

Testa `GetMetrics` com role `"ADMIN"`.  
Usa `idRole` existente no banco.  
Testa cenário válido e erro com IDs incorretos.

### `TestGetMetrics_OperatorRole`

Testa `GetMetrics` com role `"OPERADOR"` (ou outro papel não-admin).  
Também valida comportamento com parâmetros inválidos.

## 🔧 Pré-Requisitos

- Go 1.18 ou superior  
- Banco de dados configurado e com dados esperados  
- Módulo Go corretamente inicializado com `go mod init` e dependências instaladas

## 🚀 Como Executar os Testes

Para rodar todos os testes do projeto:

```bash
go test ./...
```

Para rodar apenas os testes do pacote `service`:

```bash
go test ./service
```

Para obter saídas detalhadas de log:

```bash
go test -v ./service
```

## 📌 Observações

- Certifique-se de que os IDs utilizados nos testes (`IDProject`, `idUser`, `idRole`) estejam presentes no banco.
- IDs como `99999999999` são utilizados intencionalmente para simular falhas e testar o tratamento de erros.
- As datas usadas são dinâmicas (`time.Now()`) para garantir atualidade nos testes.
- Os testes assumem que os papéis "ADMIN", "OPERADOR", etc., estão devidamente cadastrados no banco com os respectivos `idRole`.

## 📈 Sugestões Futuras

- Implementar mocks para evitar dependência direta com o banco de dados.
- Separar testes de unidade e testes de integração.
- Utilizar bibliotecas como `gomock`, `testify/mock` ou `testcontainers` para simular banco em memória ou containers.