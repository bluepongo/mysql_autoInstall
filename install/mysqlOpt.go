package install

import (
	"database/sql"
	"fmt"
	"github.com/romberli/log"

	_ "github.com/go-sql-driver/mysql"
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

// Init mysql
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

	if err := database.Ping(); err != nil {
		fmt.Println("[Warn]mysql open failed, error:", err)
		log.Warnf("Mysql open failed error: %s", err)
		return nil
	}
	fmt.Println("[Info]Mysql open successfully!")
	log.Info("Mysql open successfully!")
	return database
}

// mysql operation
func MySQLOperation(DB *sql.DB, query string, args ...interface{}) {

	// Start a transaction
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("[Warn]]tx failed.")
		log.Warn("tx failed.")
		return
	}

	// Prepare a sql
	stmt, err := tx.Prepare(query)
	if err != nil {
		fmt.Println("[Warn]]Prepare failed.")
		log.Warn("Prepare failed.")
		return
	}

	// Execute the sql sentence
	res, err := stmt.Exec(args)
	if err != nil {
		fmt.Println("[Warn]]Exec failed.")
		log.Warn("Exec failed.")
		return
	}

	// Commit the transaction
	tx.Commit()

	// Get the last id
	fmt.Println(res.LastInsertId())
	fmt.Println("[Info]Mysql exec successfully!")
	log.Info("Mysql exec successfully!")
}
