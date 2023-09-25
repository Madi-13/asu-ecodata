package domain

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBScheme       string `mapstructure:"POSTGRES_SCHEME"`

	ServerPort string `mapstructure:"SERVICE_PORT"`
	ServerIp   string `mapstructure:"SERVICE_IP"`

	AuthServiceUrl string        `mapstructure:"AUTH_SERVERICE_URL"`
	Keycloak       KeycloakStuct `mapstructure:"KEYCLOAK"`

	S3Url        string `mapstructure:"S3_URL"`
	S3AccessKey  string `mapstructure:"S3_ACCESS_KEY"`
	S3SercretKey string `mapstructure:"S3_SECRET_KEY"`
	S3Bucket     string `mapstructure:"S3_BUCKET"`
}

type KeycloakStuct struct {
	Host      string `mapstructure:"host"`
	Realm     string `mapstructure:"realm"`
	Client    string `mapstructure:"client"`
	Secret    string `mapstructure:"secret"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	PublicKey string `mapstructure:"public_key"`
}
