package install

// Batch create the folder to example
func CreateFolder(examName string) {
	Mkdir(MySQLMultiRoot + "/" + examName + "/data")
	Mkdir(MySQLMultiRoot + "/" + examName + "/tmp")
	Mkdir(MySQLMultiRoot + "/" + examName + "/sock")
	Mkdir(MySQLMultiRoot + "/" + examName + "/log")
	Mkdir(MySQLMultiRoot + "/" + examName + "/pid")
	Chmod(MySQLMultiRoot)
	Chown(GroupName, UserName, MySQLMultiRoot)
}

// Batch execute the cp command
func BatchCpBin() {
	Cp(MySQLMultiBin+"/mysql", BinPath)
	Cp(MySQLMultiBin+"/mysqld", BinPath)
	Cp(MySQLMultiBin+"/mysqld_safe", BinPath)
	Cp(MySQLMultiBin+"/mysqld_multi", BinPath)
	Cp(MySQLMultiBin+"/mysqldump", BinPath)
	Cp(MySQLMultiBin+"/mysqlbinlog", BinPath)
	Cp(MySQLMultiBin+"/mysql_config_editor", BinPath)
	Cp(MySQLMultiBin+"/my_print_defaults", BinPath)
	Cp(MySQLMultiBin+"/mysqladmin", BinPath)
}
