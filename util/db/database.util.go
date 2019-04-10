package db

import (
	"cxrus-app/util/db/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func setDbConnection() *model.DBConnection {

	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection := &model.DBConnection{}
	dbConnection.DBName = os.Getenv("db.name")
	dbConnection.DBUser = os.Getenv("db.user")
	dbConnection.DBDialect = os.Getenv("db.dialect")
	dbConnection.DBHost = os.Getenv("db.host")
	dbConnection.DBPort = os.Getenv("db.port")
	dbConnection.DBPassword = os.Getenv("db.password")

	return dbConnection

}

const (
	//PostgresDialect dialect for postgres
	PostgresDialect = "postgres"

	//MysqlDialect dialect for mysql
	MysqlDialect = "mysql"
)

func dialectInfo(dbConnection *model.DBConnection) string {
	switch dbConnection.DBDialect {
	case PostgresDialect:
		plSQLInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbConnection.DBHost, dbConnection.DBPort, dbConnection.DBUser, dbConnection.DBName, dbConnection.DBPassword)
		return plSQLInfo
	case MysqlDialect:
		mySQLInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConnection.DBUser, dbConnection.DBPassword, dbConnection.DBHost, dbConnection.DBPort, dbConnection.DBName)
		return mySQLInfo
	default:
		log.Fatal("There No Dialect for this driver")
		return ""
	}
}
