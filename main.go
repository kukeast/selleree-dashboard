package main

import (
	"main/modules/api"
	"main/modules/mysql"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	initDb()
	api.Api()
}

func initDb() {
	godotenv.Load()
	mysql.InitConnectionString(os.Getenv("DB1"), os.Getenv("DB2"))
}
