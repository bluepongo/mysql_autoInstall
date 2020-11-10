package main

import (
	"github.com/bluepongo/mysql_autoInstall/install"
)

const (
	RootPath       = "/usr/local/"
	DataDirPath    = "/usr/local/mysql/data/"
	BaseDirPath    = "/usr/local/mysql/"
	MySQLDPath     = "/usr/local/mysql/bin/mysqld"
	BinPanth       = "/usr/local/mysql/bin/"
	MySQlBinPath   = "/usr/bin/"
	MySQLErrPath   = "/usr/local/mysql/data/mysql.err"
	MySQLServePath = "/usr/local/mysql/support-files/mysql.server"
	LnMySQLServer  = "/etc/init.d/mysql"
	EtcPath        = "/etc/"

	ShareFilePath = "/mnt/hgfs/share/"

	MySQLTarName  = "mysql-5.7.31-linux-glibc2.12-x86_64.tar.gz"
	MySQLFileName = "mysql-5.7.31-linux-glibc2.12-x86_64"
	MyCnfFileName = "my.cnf"

	GroupName = "mysql"
	UserName  = "mysql"
	MySQL     = "mysql"

	// Need to change
)

func main() {
	// 1 Create user group and user
	//install.AddGroup(GroupName)
	//install.AddUser(GroupName, UserName)

	// 2 Execute the UntarGz command
	install.UnTarGz(ShareFilePath+MySQLTarName, ShareFilePath)

	install.Mv(ShareFilePath+MySQLFileName, ShareFilePath+MySQL)
	install.Cp(ShareFilePath+MySQL, RootPath+MySQL)

	// 3 Create the data directory under /usr/local/mysql
	install.Mkdir(DataDirPath)

	// 4 Change the permissions
	install.Chown(GroupName, UserName, BaseDirPath)
	install.Chmod(BaseDirPath)

	// 5 Copy the default my.cnf to the /etc/
	install.Cp(ShareFilePath+MyCnfFileName, EtcPath+MyCnfFileName)

	// 6 Compile, install, and initialize mysql
	install.InitMysqld(MySQLDPath, UserName, DataDirPath, BaseDirPath)
	//install.InitMs(MySQLDPath, UserName, DataDirPath, BaseDirPath)

	// 7 View the password
	install.Cat(MySQLErrPath)
	//fmt.Println("Please remember the above password!It's your mysql initial password, and copy to the next password.")

	// 8 Start the mysql service
	//install.ServiceStart(MySQLServePath)

	// 9 Establish the soft connection
	//install.Ln(MySQLServePath, LnMySQLServer)
	//install.Ln(MySQlBinPath+MySQL, BinPanth+MySQL)
}
