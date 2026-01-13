package mysql

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

// InitMysqlDB initialize mysql db
func InitMysqlDB(connection string) *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
	)

	log.Printf("DSN: %s\n", connection)
	// dsn = "root:@tcp(localhost:3306)/wec_engagement?parseTime=true&loc=Asia%2FJakarta"

	if os.Getenv("APP_ENV") == "production" {
		dbConn, errMysql = sql.Open(`nrmysql`, connection)
	} else {
		logrus.Infof("Connecting to MySQL DSN: %s", connection)
		dbConn, errMysql = sql.Open(`mysql`, connection)
	}
	if errMysql != nil {
		logrus.Panicf("Failed, Connection to Mysql Database: %v", errMysql)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbConn.SetMaxOpenConns(100)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbConn.SetMaxIdleConns(25)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	// Ping to test
	if err := dbConn.Ping(); err != nil {
		logrus.Panicf("mysql ping failed: %s", err.Error())
	}

	return dbConn
}
