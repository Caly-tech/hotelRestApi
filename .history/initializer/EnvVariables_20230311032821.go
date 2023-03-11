package initializer


func EnvVariables() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}