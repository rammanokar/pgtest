package cmd

import (
	"github.com/rammanokar/pgtest/config"
	"github.com/rammanokar/pgtest/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pgtest",
	Short: "A DevOps tool to test PostgreSQL connectivity",
	Long: `pgtest is a CLI tool designed for DevOps engineers to quickly test 
and verify PostgreSQL database connections. It supports both command-line 
arguments, interactive input mode, and configuration via a YAML file.

Config file structure (config.yaml):

host: localhost
port: 5432
user: postgres
password: your_password
dbname: your_database
sslmode: disable # "require" (default), "verify-full", "verify-ca", and "disable"

Place this file in the same directory as the pgtest binary or specify
the path using the --config flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.IsEmpty() {
			config.GetInteractiveInput()
		}
		db.TestConnection()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&config.Config.Host, "host", "H", "", "Database host")
	rootCmd.Flags().IntVarP(&config.Config.Port, "port", "p", 5432, "Database port")
	rootCmd.Flags().StringVarP(&config.Config.User, "user", "u", "", "Database user")
	rootCmd.Flags().StringVarP(&config.Config.Password, "password", "P", "", "Database password")
	rootCmd.Flags().StringVarP(&config.Config.DBName, "dbname", "d", "", "Database name")
	rootCmd.Flags().StringVarP(&config.Config.SSLMode, "sslmode", "s", "", "SSL mode (disable, require[default], verify-ca, verify-full)")
	rootCmd.Flags().StringVar(&config.ConfigFile, "config", "", "Path to config file")
}
