package repositories

import (
	"loginRegisterTemplate/types"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _connection *gorm.DB

func Connection() *gorm.DB {
	if _connection != nil {
		return _connection
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ")/" + db_name + "?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&types.User{})

	_connection = db

	return db
}

func Connected() bool {
	return !Connection().Config.DryRun
}
