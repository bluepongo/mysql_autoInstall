package installer

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

func Execute() {
	// Initialize a logger
	fileName := LogFilePath
	_, _, err := log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		panic(err)
	}
	log.Info("Initial log file success.")
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
		fmt.Println("=========Prepare to generate mycnf=========")
		// 生成对应的mycnf文件
		err = conf.GenerateMyCnf(CurrentIP, IpPort.Port)
		if err != nil {
			log.Warnf("%v", err)
			return
		}
		fmt.Println("=========Prepare to create ssh connection=========")
		install.InstallMysqlSSH(CurrentIP, IpPort.Port)
		fmt.Println("=========Finish=========")
	}
}
