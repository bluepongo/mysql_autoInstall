package dataRelated

import (
	"fmt"
	"github.com/bluepongo/mysql_autoInstall/install"
)

func CreateTest() {

	// Start the Mysql operations
	// Initialize the database
	install.InitDB("root", "", "3306", "mysql")
	// Create a admin user
	install.ExecMysql("grant shutdown on *.* to 'admin'@'localhost' identified by 'admin'")
	// Create a replica account
	install.ExecMysql("grant replication slave, replication client on *.* to 'replication'@'%' identified by 'admin'")
	// Create a database and use it
	install.ExecMysql("drop database spdb")
	install.ExecMysql("create database spdb")
	install.ExecMysql("use spdb")
	// Create the test data
	install.ExecMysql(fmt.Sprintf(
		`create table t01(
			id int(10) primary key auto_increment comment 'pid',
			name varchar(100) comment 'name'
		) engine= innodb default charset=utf8mb4`))
	install.InsertMysql("insert into t01(name) values('a'), ('b'), ('c')")
	install.QueryMany("select * from t01")
}
