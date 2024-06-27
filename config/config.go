package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var Config DatabaseConfig
var ConfigFile string

func LoadConfig() error {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return nil
}

func IsEmpty() bool {
	return Config.Host == "" || Config.Port == 0 || Config.User == "" || Config.DBName == ""
}

func GetInteractiveInput() {
	reader := bufio.NewReader(os.Stdin)

	if Config.Host == "" {
		Config.Host = prompt(reader, "Enter database host", "localhost")
	}

	if Config.Port == 0 {
		portStr := prompt(reader, "Enter database port", "5432")
		port, err := strconv.Atoi(portStr)
		if err != nil {
			color.Yellow("Invalid port number. Using default: 5432")
			Config.Port = 5432
		} else {
			Config.Port = port
		}
	}

	if Config.User == "" {
		Config.User = prompt(reader, "Enter database user", "postgres")
	}

	if Config.Password == "" {
		Config.Password = promptPassword(reader, "Enter database password")
	}

	if Config.DBName == "" {
		Config.DBName = prompt(reader, "Enter database name", "postgres")
	}

	if Config.SSLMode == "" {
		Config.SSLMode = prompt(reader, "Enter SSL mode (disable, require, verify-ca, verify-full)", "disable")
	}
}

func prompt(reader *bufio.Reader, promptText, defaultValue string) string {
	color.Blue("%s [%s]: ", promptText, defaultValue)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}

func promptPassword(reader *bufio.Reader, promptText string) string {
	color.Blue(promptText + ": ")
	password, _ := reader.ReadString('\n')
	return strings.TrimSpace(password)
}
