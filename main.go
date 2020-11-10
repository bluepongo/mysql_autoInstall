package main

import "github.com/bluepongo/mysql_autoInstall/install"

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
	//install.Mv(ShareFilePath+MySQLFileName, RootPath)
	//install.Mv(RootPath+MySQLFileName, BaseDirPath)
}