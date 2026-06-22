package configuration

import "os" // permet de changer la configuration sans recomplier le code en stockant sur des variables d'environnement

type Config struct {
	APIKey  string
	Port    string
	MongoURI string
	MongoDB  string
}

func LoadConfig() Config {
	return Config{
		APIKey:  getEnv("API_KEY", "humancraft"),
		Port:    getEnv("PORT", "8080"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:  getEnv("MONGO_DB", "todolist"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
