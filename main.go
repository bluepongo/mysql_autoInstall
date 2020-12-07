package main

import (
	"flag"
	"fmt"
	"github.com/bluepongo/mysql_autoInstall/conf"
	"github.com/bluepongo/mysql_autoInstall/install"
	"github.com/bluepongo/mysql_autoInstall/parameters"
	"github.com/romberli/log"
	"strconv"
)

const (
	LogFilePath = "/tmp/test.log"
	// Need to change
)

func main() {

	// Initialize a logger
	fileName := LogFilePath
	_, _, err := log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		panic(err)
	}
	log.Info("Initial log file success.")

	// Install mysql multi
	//fmt.Println("=========Start install the MySQL5.7=========")
	//install.InstallMySQLMul()
	//fmt.Println("=========Install the MySQL5.7 success=========")
	//
	//// Start the Mysql operations
	//fmt.Println("=========Start the Mysql operations=========")
	//dataRelated.CreateTest()
	//
	//// Build master-slave replication
	//fmt.Println("=========Start build master-slave replication=========")
	//install.BuildMS()
	//fmt.Println("=========Setup master slave copied successfully=========")

	// 接受参数
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	// 解析传入参数
	IpPorts := parameters.ExtractIP(*parameters.Ip)
	for _, IpPort := range IpPorts {
		CurrentIP := IpPort.Ip
		CurrentPort, _ := strconv.Atoi(IpPort.Port)
		log.Infof("Current:Ip:%s, Port:%d", CurrentIP, CurrentPort)
		// Install mysql remotely via SSH connection
		fmt.Println("=========Prepare to create ssh connection=========")
		// 生成对应的mycnf文件
		err = conf.GenerateMyCnf("192.168.59.02", "3306")
		if err != nil {
			log.Warnf("%v", err)
			return
		}

		install.InstallMysqlSSH(CurrentIP)

	}

}
