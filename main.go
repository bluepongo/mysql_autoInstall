package main

import (
	"fmt"
	dataRelated "github.com/bluepongo/mysql_autoInstall/data_related"
	"github.com/bluepongo/mysql_autoInstall/install"
)

const (
// Need to change
)

func main() {
	// Install mysql
	//install.InstallMysql()

	// Install mysql multi
	fmt.Println("=========Start install the MySQL5.7=========")
	install.InstallMySQLMul()

	// Start the Mysql operations
	fmt.Println("=========Start the Mysql operations=========")
	dataRelated.CreateTest()

	// Build master-slave replication
	fmt.Println("=========Start build master-slave replication=========")
	install.BuildMS()
	fmt.Println("=========Setup master slave copied successfully=========")

}
