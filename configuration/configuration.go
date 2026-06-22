package configuration

import "os" // permet de changer la configuration sans recomplier le code en stockant sur des variables d'environnement

type Config struct {
	APIKey string
	DBPath string
	Port   string
}

func LoadConfig() Config {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "humancraft"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "todos.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		APIKey: apiKey,
		DBPath: dbPath,
		Port:   port,
	}
}
