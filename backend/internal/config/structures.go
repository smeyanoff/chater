package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Auth     AuthConfig
	CORS     CORSConfig
}

type AppConfig struct {
	AppPort  string
	AppHost  string
	AppAdmin string
	AppPwd   string
}

type DatabaseConfig struct {
	DBPort  int
	DBHost  string
	DBTable string
	DBUser  string
	DBPwd   string
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
