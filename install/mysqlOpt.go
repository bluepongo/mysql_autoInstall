package install

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize a database function
func InitDB(Username, Password, Port, Database string) (err error) {
	// DSN:Data Source Name
	dsn := Username + ":" + Password + "@tcp(localhost:" + Port + ")/" + Database + "?charset=utf8"
	// Check the password
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Exec failed, err:%v\n", err)
		return err
	}
	// attempt connect to the databse
	err = db.Ping()
	if err != nil {
		fmt.Printf("Initialize failed, err:%v\n", err)
		return err
	}
	fmt.Println("Database initialization succeeded.")
	return nil
}

// Exec the sql string
func ExecMysql(sqlStr string) {
	_, err := db.Exec(sqlStr)
	fmt.Println("mysql >", sqlStr)
	if err != nil {
		fmt.Printf("Exec failed, err:%v\n", err)
		return
	}
	fmt.Println("Exec success.")
}

// Exec the insert sql string
func InsertMysql(sqlStr string) {
	ret, err := db.Exec(sqlStr)
	fmt.Println("mysql >", sqlStr)
	if err != nil {
		fmt.Printf("Insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // the new data id
	if err != nil {
		fmt.Printf("Get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("Insert success, the id is %d.\n", theID)
}

type test struct {
	id   int
	name string
}

// Query many rows
func QueryMany(sqlStr string) {
	rows, err := db.Query(sqlStr)
	fmt.Println("mysql >", sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var t test
		err := rows.Scan(&t.id, &t.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d  name:%s\n", t.id, t.name)
	}
}
