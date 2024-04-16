package config

import (
	"os"
	"strconv"
)

var (
	ApiPort          int
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDataBase string
	IPFSHost         string
)

func LoadConfig() {
	var err error
	if ApiPort, err = strconv.Atoi(os.Getenv("API_PORT")); err != nil {
		panic(err)
	}
	PostgresHost = os.Getenv("POSTGRES_HOST")
	if PostgresPort, err = strconv.Atoi(os.Getenv("POSTGRES_PORT")); err != nil {
		panic(err)
	}
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresDataBase = os.Getenv("POSTGRES_DB")

	IPFSHost = os.Getenv("IPFS_HOST")
}
