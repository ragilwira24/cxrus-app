package db

import (
	appModel "cxrus-app/service/models"
	"cxrus-app/util/db/model"
	"cxrus-app/util/db/mysql"
	"cxrus-app/util/db/postgres"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"

	// Import the dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbs *gorm.DB

//InitDB to initialize DB
func InitDB() {

	dbCon := setDbConnection()
	dialectInfo := dialectInfo(dbCon)
	log.Println(dbCon.DBDialect + " configuration :" + dialectInfo)

	db, err := DialectValidation(dbCon, dialectInfo)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	dbs = db

	for _, obj := range []interface{}{
		&appModel.User{},
	} {
		if err := db.AutoMigrate(obj); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Auto Migrating", reflect.TypeOf(obj).Name(), "...")
		}
	}

	log.Println("Successfully connected")
}

// DialectValidation initialize based on dialect
func DialectValidation(dbCon *model.DBConnection, dialectInfo string) (*gorm.DB, error) {

	switch dbCon.DBDialect {
	case PostgresDialect:
		log.Println("Initializing Postgres Dialect")
		return postgres.InitDBPostgres(dbCon, dialectInfo)
	case MysqlDialect:
		log.Println("Initializing Mysql Dialect")
		return mysql.InitDBMysql(dbCon, dialectInfo)
	default:
		log.Fatal("No Dialect to Initialize")
		return nil, errors.New("No Dialect to Initialize")

	}
}

// GetDB to get DB
func GetDB() *gorm.DB {
	return dbs
}
