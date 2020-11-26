package install

const (
	RootPath     = "/usr/local/"
	DataDirPath  = "/usr/local/mysql/data/"
	BaseDirPath  = "/usr/local/mysql/"
	MySQLDPath   = "/usr/local/mysql/bin/mysqld"
	EtcPath      = "/etc/"
	LimitsPath   = "/etc/security/"
	BashFilePath = "/home/mysql/"

	ShareFilePath = "/mnt/hgfs/share/"

	MySQLTarName  = "mysql-5.7.31-linux-glibc2.12-x86_64.tar.gz"
	MySQLFileName = "mysql-5.7.31-linux-glibc2.12-x86_64"
	MyCnfFileName = "my.cnf"
	LimitsFile    = "limits.conf"
	BashFile      = ".bash_profile"
	AutoCnfFile   = "auto.cnf"

	GroupName = "mysql"
	UserName  = "mysql"
	MySQL     = "mysql"

	// MySQLMulti
	MySQLMulti_3306       = "/usr/local/mysql/3306/"
	MySQLMulti_3307       = "/usr/local/mysql/3307"
	MySQLMultiBinlog_3306 = "/usr/local/mysql/3306/binlog"
	MySQLMultiData_3306   = "/usr/local/mysql/3306/data"
	MySQLMultiData_3307   = "/usr/local/mysql/3307/data/"
	MySQLMultiRoot        = "/usr/local/mysql"
	MySQLMultiBin         = "/usr/local/mysql/bin"
	MultiDataDir          = "/usr/local/mysql/3306/data"

	RelatedPath = "./related/"

	BinPath = "/usr/bin"

	PortNum_3306 = "3306"
	PortNum_3307 = "3307"

	// Need to change

)

func InstallMysql() {
	// 1 Create user group and user
	AddGroup(GroupName)
	AddUser(GroupName, UserName)

	// 2 Execute the UntarGz command
	UnTarGz(ShareFilePath+MySQLTarName, ShareFilePath)

	Mv(ShareFilePath+MySQLFileName, ShareFilePath+MySQL)
	Cp(ShareFilePath+MySQL, RootPath+MySQL)

	// 3 Create the data directory under /usr/local/mysql
	Mkdir(DataDirPath)

	// 4 Change the permissions
	Chown(GroupName, UserName, BaseDirPath)
	Chmod(BaseDirPath)

	// 5 Copy the default my(basic).cnf to the /etc/
	Cp(ShareFilePath+MyCnfFileName, EtcPath+MyCnfFileName)

	// 6 Compile, install, and initialize mysql
	InitMysql(MySQLDPath, UserName, DataDirPath, BaseDirPath)
}

// Install multiple instances of mysql
func InstallMySQLMul() {
	//// 1、Create mysql user and user group
	//AddGroup(GroupName)
	//AddUser(GroupName, UserName)

	// 2、Create the catalogue
	Mkdir(MySQLMultiBinlog_3306)
	Mkdir(MySQLMultiData_3306)
	Chown(GroupName, UserName, MySQLMultiRoot)

	// 3、Alter the file limits.cnf
	Cp(RelatedPath+LimitsFile, LimitsPath+LimitsFile)

	// 4、Unzip the file
	UnTarGz(RelatedPath+MySQLTarName, RelatedPath)

	// 5、Move the mysql file
	Mv(RelatedPath+MySQLFileName, RelatedPath+MySQL)
	Cp(RelatedPath+MySQL+"/*", MySQLMultiRoot)

	// 6、Alter the my.cnf file
	Cp(RelatedPath+MyCnfFileName, EtcPath+MyCnfFileName)

	// 7、Create the folder data/tmp/log/pid/sock to exam
	CreateFolder(PortNum_3306)

	// 8、Batch execute cp command
	BatchCpBin()

	// 9、Alter the .bash_profile
	Cp(RelatedPath+BashFile, BashFilePath+BashFile)

	// 10、Initializes the mysql instance
	MultiInitMysql(MySQLDPath, UserName, BaseDirPath, MultiDataDir)

	// 11、Start the example
	MultiStartMysql(PortNum_3306)
}

// Build master-slave replication
func BuildMS() {
	// 1、Stop the 3306 instance
	MultiStopMysql(PortNum_3306)

	// 2、Copy data file from 3306 to 3307
	Cp(MySQLMulti_3306, MySQLMulti_3307)
	Chown(GroupName, UserName, MySQLMulti_3307)
	Rm(MySQLMultiData_3307 + AutoCnfFile)

	// 3、Start the 3306、3307
	MultiStartMysql(PortNum_3306)
	MultiStartMysql(PortNum_3307)
	// Initialize the database
	InitDB("root", "", "3307", "mysql")
	// change the master
	ExecMysql("change master to master_host='localhost', master_port=3306, master_user='replication', master_password='admin', master_auto_position=1")
	ExecMysql("start slave")

}
