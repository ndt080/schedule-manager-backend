package configs

type SmtpConfig struct {
	Username       string `envconfig:"SMTP_USERNAME"`
	Password       string `envconfig:"SMTP_PASSWORD"`
	Host           string `envconfig:"SMTP_HOST"`
	Port           string `envconfig:"SMTP_PORT"`
	ServerBasePath string
}
