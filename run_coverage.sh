#!/bin/bash

# Define o diretório onde o relatório de cobertura será salvo (dentro de src)
coverageDir="./src/coverage_report"

# Cria o diretório coverage_report se não existir
mkdir -p "$coverageDir"
echo "Diretório '$coverageDir' criado."

# Exclui os diretórios dto, models e os arquivos *_test.go da pesquisa
PACKAGES=$(go list ./... | grep -v '/dto' | grep -v '/models' | grep -v '_test.go')

# Executa os testes e gera o arquivo de cobertura no diretório src/coverage_report
echo "Executando testes e gerando cobertura..."
go test -v -race -coverprofile="$coverageDir/coverage.out" -covermode=atomic $PACKAGES || true

# Verifica se o arquivo de cobertura foi gerado
if [ -f "$coverageDir/coverage.out" ]; then
    echo "Arquivo de cobertura gerado em $coverageDir/coverage.out"
else
    echo "⚠️ O arquivo de cobertura não foi gerado."
    exit 1
fi

# Executa a conversão do arquivo .out para .html para melhor visualização
echo "Gerando relatório HTML de cobertura..."
go tool cover -html="$coverageDir/coverage.out" -o "$coverageDir/coverage.html"

# Verifica se o HTML foi gerado
if [ -f "$coverageDir/coverage.html" ]; then
    echo "✅ Relatório HTML gerado em $coverageDir/coverage.html"
else
    echo "⚠️ Falha ao gerar o arquivo HTML. Verifique se há erros durante a conversão."
    exit 1
fi

# Como usar:
# Salvar o script: Salve o script em um arquivo chamado, por exemplo, generate_coverage_report.sh.
# Tornar o script executável: No terminal, torne o script executável com o comando: chmod +x generate_coverage_report.sh
# Executar o script: Execute o script com o seguinte comando: ./generate_coverage_report.sh