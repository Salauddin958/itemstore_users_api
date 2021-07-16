package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	logger "github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
	mysqlUsersPort     = "mysql_users_port"
)

var (
	Client *sql.DB
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		os.Getenv(mysqlUsersUsername),
		os.Getenv(mysqlUsersPassword),
		os.Getenv(mysqlUsersHost),
		os.Getenv(mysqlUsersPort),
		os.Getenv(mysqlUsersSchema),
	)
	fmt.Println("dataSourceName :: ", dataSourceName)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
