package install

import (
	"database/sql"
	"fmt"
	"github.com/romberli/log"
)

const (
	Engine = "mysql"
	User   = "root"
	Pass   = ""
	Ip     = "127.0.0.1"
	Port   = "3306"
	Table  = "test"
)

type DbInfo struct {
	Engine string
	User   string
	Pass   string
	Ip     string
	Port   string
	Table  string
}

func MySQLInit() *sql.DB {

	db1 := DbInfo{
		Engine,
		User,
		Pass,
		Ip,
		Port,
		Table,
	}

	// Initialize a logger
	fileName := LogFilePath
	_, _, err := log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		fmt.Printf("Init logger failed.\n%s", err.Error())
	}

	database, err := sql.Open(db1.Engine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf-8", db1.User, db1.Pass, db1.Ip, db1.Port, db1.Table))

	if err != nil {
		fmt.Println("mysql open failed, error:", err)
		log.Warnf("Mysql open failed error: %s", err)
		return nil
	}
	fmt.Println("Mysql open successfully!")
	log.Info("Mysql open successfully!")
	return database
}
