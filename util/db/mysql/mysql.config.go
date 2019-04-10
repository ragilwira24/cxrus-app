package mysql

import (
	"cxrus-app/util/db/model"

	"github.com/jinzhu/gorm"
	// Import the dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitDBMysql for Postgre
func InitDBMysql(dbCon *model.DBConnection, dialectInfo string) (*gorm.DB, error) {
	db, err := gorm.Open(dbCon.DBDialect, dialectInfo)
	return db, err
}
