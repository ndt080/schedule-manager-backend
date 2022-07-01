package configs

import (
    "os"
    "strconv"
    "time"
)

type HttpConfig struct {
    Port           int
    MaxHeaderBytes int
    ReadTimeout    time.Duration
    WriteTimeout   time.Duration
}

type AuthConfig struct {
    SigningKey      string
    AccessTokenTTL  time.Duration
    RefreshTokenTTL time.Duration
}

type ServerConfig struct {
    Http HttpConfig
    Auth AuthConfig
}

func NewServerConfig(configDir string) *ServerConfig {
    if err := ParseConfigFile(configDir); err != nil {
        return nil
    }

    config := ServerConfig{}
    if err := Unmarshal(&config); err != nil {
        return nil
    }
    return &config
}

func GetServerPort(config *ServerConfig, envName string) string {
    if envValue := os.Getenv(envName); envValue != "" {
        return envValue
    }

    return strconv.Itoa(config.Http.Port)
}
