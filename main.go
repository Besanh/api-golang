package main

import (
	"api-golang/api"
	apiv1 "api-golang/api/v1"
	"api-golang/repository"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func init() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?readTimeout=%ds&writeTimeout=%ds", "root", "TEL4VN!@#Intern2021@12", "113.164.246.25", 4308, "training-api", 10, 30)
	sqldb, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	if err := db.Ping(); err != nil {
		panic(err)
	}
	repository.Db = db
}

func main() {
	server := api.DeclareServer()

	// Ben repo la khai bao
	// Khoi tao repo
	repository.UserRepo = repository.NewUserRepo()
	apiv1.APIUser(server.Engine)
	server.Run()
}
