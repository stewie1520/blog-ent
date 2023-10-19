package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/viper"
)

// C is the global config.
var _c *Config

func C() *Config {
	return _c
}

type Config struct {
	Env           string `mapstructure:"ENV"`
	Port          int16  `mapstructure:"PORT"`
	GrpcPort      int16  `mapstructure:"GRPC_PORT"`
	ApiDomain     string `mapstructure:"API_DOMAIN"`
	WebsiteDomain string `mapstructure:"WEBSITE_DOMAIN"`
	GinMode       string `mapstructure:"GIN_MODE"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	Token         struct {
		PrivateKey      string `mapstructure:"PRIVATE_KEY"`
		PublicKey       string `mapstructure:"PUBLIC_KEY"`
		AccessTokenTTL  int    `mapstructure:"ACCESS_TOKEN_TTL"`
		RefreshTokenTTL int    `mapstructure:"REFRESH_TOKEN_TTL"`
	} `mapstructure:"TOKEN"`
}

func (c *Config) Validate() error {
	err := validation.ValidateStruct(c,
		validation.Field(&c.Port, validation.Required),
		validation.Field(&c.DatabaseURL, validation.Required, is.URL),
	)

	if err != nil {
		return err
	}

	return validation.ValidateStruct(&c.Token,
		validation.Field(&c.Token.PrivateKey, validation.Required),
		validation.Field(&c.Token.PublicKey, validation.Required),
		validation.Field(&c.Token.AccessTokenTTL, validation.Required),
		validation.Field(&c.Token.RefreshTokenTTL, validation.Required),
	)
}

func (c *Config) IsProd() bool {
	return c.Env == "production"
}

func Init() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	err = config.Validate()
	if err != nil {
		return nil, err
	}

	_c = config

	return config, nil
}
