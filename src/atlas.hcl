env "gorm" {
  driver = "go://github.com/iNineBD/Track-5Sem2025Server/src/pkg/models.GetModels"
  url = "postgres://eduardo:a9c3t1@144.22.212.19:5432/track?sslmode=disable"
  dev = "docker://postgres/14/dev?search_path=public"
  
  migration {
    dir = "file://migrations"
  }
}