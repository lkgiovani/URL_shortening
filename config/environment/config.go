package environment

import (
	"url_shortening/pkg/env"
	"url_shortening/pkg/projectError"
)

type Config struct {
	HTTP struct {
		Url  string
		Port int
	}
	DB struct {
		DataSource string
	}
	URL_SHORTENED_PREFIX string
	REDIS                struct {
		Address string
	}
	JWT_SECRET string
}

func NewConfig() (*Config, error) {

	httpUrl, err := getString("URL", "Error loading HTTP URL")
	if err != nil {
		return nil, err
	}

	httpPort, err := getInt("PORT", "Error loading HTTP Port")
	if err != nil {
		return nil, err
	}

	dbDataSource, err := getString("DB_DATA_SOURCE", "Error loading DB Data Source")
	if err != nil {
		return nil, err
	}

	urlShortenedPrefix, err := getString("URL_SHORTENED_PREFIX", "Error loading URL Shortened Prefix")
	if err != nil {
		return nil, err
	}

	redisAddress, err := getString("REDIS_ADDRESS", "Error loading Redis Address")
	if err != nil {
		return nil, err
	}

	jwtSecret, err := getString("JWT_SECRET", "Error loading JWT Secret")
	if err != nil {
		return nil, err
	}

	return &Config{
		HTTP: struct {
			Url  string
			Port int
		}{
			Url:  httpUrl,
			Port: httpPort,
		},
		DB: struct {
			DataSource string
		}{
			DataSource: dbDataSource,
		},
		URL_SHORTENED_PREFIX: urlShortenedPrefix,
		REDIS: struct {
			Address string
		}{
			Address: redisAddress,
		},
		JWT_SECRET: jwtSecret,
	}, nil
}

func getInt(key, errorMessage string) (int, error) {
	value, err := env.GetEnvOrDieAsInt(key)

	if err != nil {
		return 0, &projectError.Error{
			Code:    projectError.EINVALID,
			Message: errorMessage,
		}
	}

	return value, nil

}

func getString(key, errorMessage string) (string, error) {
	value, err := env.GetEnvOrDie(key)
	if err != nil {
		return "", &projectError.Error{
			Code:    projectError.EINVALID,
			Message: errorMessage,
		}
	}
	return value, nil
}
