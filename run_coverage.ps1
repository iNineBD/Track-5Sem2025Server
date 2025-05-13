# Navega até a pasta src (onde está o arquivo go.mod)
Set-Location -Path ".\src"

# Define o diretório onde o relatório de cobertura será salvo
$coverageDir = ".\coverage_report"

# Cria o diretório se não existir
if (-not (Test-Path -Path $coverageDir)) {
    New-Item -ItemType Directory -Force -Path $coverageDir
}

# Define o caminho para o arquivo de cobertura (coverage.out) dentro de coverage_report
$coverageOutPath = "$coverageDir\coverage.out"

# Pega todos pacotes dentro de ./src, ignorando /dto e /models
$packages = go list ./... | Where-Object {$_ -notmatch "/dto" -and $_ -notmatch "/models"}

# Se houver pacotes, executa os testes
if ($packages.Count -gt 0) {
    # Executa os testes e gera o arquivo de cobertura no caminho especificado
    Write-Host "Executando testes e gerando cobertura..."
    go test -v -covermode=atomic $packages -coverprofile="$coverageOutPath"
    
    # Verifica se o arquivo de cobertura foi gerado
    if (Test-Path $coverageOutPath) {
        Write-Host "Arquivo de cobertura gerado em $coverageOutPath" -ForegroundColor Green
    } else {
        Write-Host "⚠️ O arquivo de cobertura não foi gerado." -ForegroundColor Red
    }
} else {
    Write-Error "Nenhum pacote válido encontrado para teste."
}

# Verifica se o arquivo de cobertura foi gerado e contém dados
if (Test-Path $coverageOutPath) {
    $coverageData = Get-Content $coverageOutPath
    if ($coverageData.Length -gt 0) {
        Write-Host "Arquivo coverage.out encontrado e contém dados! Iniciando a conversão para HTML..." -ForegroundColor Green
        
        # Define o caminho para o arquivo HTML
        $htmlCoveragePath = "$coverageDir\coverage.html"

        # Converte o arquivo .out para .html
        go tool cover -html="$coverageOutPath" -o "$htmlCoveragePath"

        # Verifica se o HTML foi gerado
        if (Test-Path $htmlCoveragePath) {
            Write-Host "✅ Relatório HTML gerado em $htmlCoveragePath" -ForegroundColor Green
        } else {
            Write-Host "⚠️ Falha ao gerar o arquivo HTML. Verifique se há erros durante a conversão." -ForegroundColor Red
        }
    } else {
        Write-Host "⚠️ O arquivo coverage.out está vazio. Verifique os testes." -ForegroundColor Red
    }
} else {
    Write-Host "⚠️ Arquivo de cobertura não gerado. Verifique os testes." -ForegroundColor Red
}
