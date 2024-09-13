package config

type Config struct {
	App  AppConfig
	DB   DatabaseConfig
	Auth AuthConfig
	CORS CORSConfig
}

type AppConfig struct {
	Port  string
	Host  string
	Admin string
	Pwd   string
}

type DatabaseConfig struct {
	Port  int
	Host  string
	Table string
	User  string
	Pwd   string
}

type AuthConfig struct {
	JWTKey                string
	JWTKeyExpirationTimeH int
	OpenAIKey             string
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}
