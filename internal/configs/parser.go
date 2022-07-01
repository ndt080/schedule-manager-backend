package configs

import "github.com/spf13/viper"

func Unmarshal(config *ServerConfig) error {
	if err := viper.UnmarshalKey("http.port", &config.Http.Port); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http.maxHeaderBytes", &config.Http.MaxHeaderBytes); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http.readTimeout", &config.Http.ReadTimeout); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http.writeTimeout", &config.Http.WriteTimeout); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth.signingKey", &config.Auth.SigningKey); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth.accessTokenTTL", &config.Auth.AccessTokenTTL); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth.refreshTokenTTL", &config.Auth.RefreshTokenTTL); err != nil {
		return err
	}

	return nil
}

func ParseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.MergeInConfig()
}
