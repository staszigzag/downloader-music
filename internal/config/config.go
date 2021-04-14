package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type BotConfig struct {
	TelegramToken string `mapstructure:"telegram_token"`
	SudoChatId    int64  `mapstructure:"sudo_chat_id"`
	Messages      Messages
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start  string
	Finish string
	Help   string
}

type Errors struct {
	Default        string
	UnknownCommand string
}

type HTTPConfig struct {
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

type FileStorageConfig struct {
	Path   string
	Bucket string
}

type DbConfig struct {
	DbPassword string `mapstructure:"db_password"`
}

type Config struct {
	Bot         BotConfig
	HTTP        HTTPConfig
	DB          DbConfig
	FileStorage FileStorageConfig
	Debug       bool
}

// Init populates Config struct with values from config file
// located at filepath and environment variables
func Init(path string) (*Config, error) {
	if err := parseConfigFile(path); err != nil {
		return nil, err
	}
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := bindEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	fmt.Println(cfg)

	return &cfg, nil
}

func parseConfigFile(filePath string) error {
	sep := "/"
	path := strings.Split(filePath, sep)
	folder := strings.Join(path[:len(path)-1], sep)
	fileName := path[len(path)-1]

	viper.AddConfigPath(folder)   // folder
	viper.SetConfigName(fileName) // config file name

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Bot.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Bot.Messages.Errors); err != nil {
		return err
	}

	// Unmarshal for environment variable
	{
		if err := viper.Unmarshal(&cfg.Bot); err != nil {
			return err
		}

		if err := viper.Unmarshal(&cfg.DB); err != nil {
			return err
		}
	}

	return nil
}

func bindEnv() error {
	// viper.SetEnvPrefix("db")
	if err := viper.BindEnv("db_password"); err != nil {
		return err
	}
	if err := viper.BindEnv("telegram_token"); err != nil {
		return err
	}
	if err := viper.BindEnv("sudo_chat_id"); err != nil {
		return err
	}

	return nil
}
