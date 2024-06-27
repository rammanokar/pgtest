package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"github.com/rammanokar/pgtest/config"
)

func TestConnection() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Config.Host, config.Config.Port, config.Config.User,
		config.Config.Password, config.Config.DBName, config.Config.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		color.Red("Failed to open a DB connection: %v", err)
		printTroubleshootingInfo(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		color.Red("Failed to ping the database: %v", err)
		printTroubleshootingInfo(err)
		return
	}

	color.Green("Successfully connected to the database!")

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		color.Red("Failed to execute query: %v", err)
		return
	}

	color.Cyan("PostgreSQL version: %s", version)
}

func printTroubleshootingInfo(err error) {
	fmt.Println("\nTroubleshooting Information:")

	if strings.Contains(err.Error(), "no pg_hba.conf entry") {
		color.Yellow("This error suggests that the PostgreSQL server is not configured to accept connections from your client's IP address.")
		fmt.Println("Possible solutions:")
		fmt.Println("1. Check the pg_hba.conf file on the PostgreSQL server and ensure it allows connections from your client's IP.")
		fmt.Println("2. If you're connecting remotely, make sure the PostgreSQL server is configured to accept remote connections.")
		fmt.Println("3. Verify that the server's firewall allows incoming connections on the PostgreSQL port (usually 5432).")
	} else if strings.Contains(err.Error(), "connection refused") {
		color.Yellow("The connection was refused. This could mean the server is not running or is not accessible.")
		fmt.Println("Possible solutions:")
		fmt.Println("1. Verify that the PostgreSQL server is running.")
		fmt.Println("2. Check if the host and port are correct.")
		fmt.Println("3. Ensure that the server's firewall is not blocking the connection.")
	} else if strings.Contains(err.Error(), "password authentication failed") {
		color.Yellow("The provided username or password is incorrect.")
		fmt.Println("Possible solutions:")
		fmt.Println("1. Double-check the username and password.")
		fmt.Println("2. Ensure the user has the necessary permissions to connect to the specified database.")
	}

	fmt.Println("\nConnection details:")
	fmt.Printf("Host: %s\n", config.Config.Host)
	fmt.Printf("Port: %d\n", config.Config.Port)
	fmt.Printf("User: %s\n", config.Config.User)
	fmt.Printf("Database: %s\n", config.Config.DBName)
	fmt.Printf("SSL Mode: %s\n", config.Config.SSLMode)
}
