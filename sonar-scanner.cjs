// sonar-scanner.cjs
const scanner = require('sonar-scanner');

scanner({
  serverUrl: 'http://144.22.212.19:9001',
  token: process.env.SONAR_TOKEN,
  options: {
    'sonar.projectKey': 'Track-5Sem2025Server',
    'sonar.projectName': 'Track-5Sem2025Server',
    'sonar.projectVersion': '1.0',
    'sonar.sources': 'src',
    'sonar.sourceEncoding': 'UTF-8',
    'sonar.go.coverage.reportPaths': 'src/coverage.out',
    'sonar.exclusions': '**/*_test.go,**/models/**,**/dto/**,**/docs/**',
    'sonar.coverage.exclusions': '**/*_test.go,**/models/**,**/dto/**,**/docs/**',
    'sonar.branch.name': process.env.GITHUB_REF_NAME || 'main',
    'sonar.verbose': 'true'
  }
});