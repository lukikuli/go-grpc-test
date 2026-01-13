package mysql

import (
	"fmt"
	"go-grpc/pkg/dotenv"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MysqlDbConf struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func (m *MysqlDbConf) SetHost(host string) {
	m.host = host
}

func (m *MysqlDbConf) GetHost() string {
	return m.host
}

func (m *MysqlDbConf) SetPort(port string) {
	m.port = port
}

func (m *MysqlDbConf) GetPort() string {
	return m.port
}

func (m *MysqlDbConf) SetUser(user string) {
	m.user = user
}

func (m *MysqlDbConf) GetUser() string {
	return m.user
}

func (m *MysqlDbConf) SetPassword(password string) {
	m.password = password
}

func (m *MysqlDbConf) GetPassword() string {
	return m.password
}

func (m *MysqlDbConf) SetDatabase(database string) {
	m.database = database
}

func (m *MysqlDbConf) GetDatabase() string {
	return m.database
}

func (m *MysqlDbConf) FormatDSN() (string, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	conn := mysql.Config{
		User:                 m.user,
		Passwd:               m.password,
		DBName:               m.database,
		Addr:                 fmt.Sprintf("%s:%s", m.host, m.port),
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  location,
		AllowNativePasswords: true,
	}

	return conn.FormatDSN(), nil
}

// MYSQLCONFIG get default mysql connection config
func MYSQLCONFIG() string {
	mysqlConf := MysqlDbConf{}
	mysqlConf.SetHost(dotenv.GetString("DB_HOST", "localhost"))
	mysqlConf.SetPort(dotenv.GetString("DB_PORT", "3306"))
	mysqlConf.SetUser(dotenv.GetString("DB_USER", "root"))
	mysqlConf.SetPassword(dotenv.GetString("DB_PASS", ""))
	mysqlConf.SetDatabase(dotenv.GetString("DB_NAME", "grpc"))

	mysqlDsn, err := mysqlConf.FormatDSN()
	if err != nil {
		log.Panic(err.Error())
	}

	return mysqlDsn
}
