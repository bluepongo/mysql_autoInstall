package install

import (
	"github.com/romberli/log"
)

const (
	HostIP      = "192.168.59.2"
	PortNum     = 22
	SSHUserName = "root"
	SSHPassWord = "root"
	SSHDatadir  = "/mysqldata/mysql3306/data"
	SSHLog      = "/mysqldata/mysql3306/log"
)

// Install mysql remotely via SSH connection
func InstallMysqlSSH(ip string) {

	// Establish the ssh connection
	log.Info("==========Install mysql remotely started==========")
	sshConn, err := EstablishSSHConnect(ip, PortNum, SSHUserName, SSHPassWord)
	if err != nil {
		log.Warnf("Can't establish the ssh connection: %v", err)
		return
	}
	log.Info("==========Install mysql remotely completed==========")

	// Execute remote shell command
	// Create the user and group
	result, stdOut, err := AddUserGroupSSH(sshConn, UserName, GroupName)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	log.Info("==========Add user and group completed==========")

	// UnTarGz the file
	stdErr, err := UnTarLocal(RelatedPath, MySQLTarName, MySQLFileName, MySQL)
	if err != nil {
		log.Warnf("%v: %s", err, stdErr)
	}

	// Move the folder to remote
	err = CopyMysqlToRemote(sshConn, RelatedPath, RootPath)
	if err != nil {
		log.Warnf("Can't copy the file to remote: %v", err)
	}
	err = CopyMysqlToRemote(sshConn, RelatedPath+MyCnfFileName, EtcPath)
	if err != nil {
		log.Warnf("Can't copy the file to remote: %v", err)
	}

	result, stdOut, err = MkdirSSH(sshConn, SSHDatadir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	result, stdOut, err = ChownSSH(sshConn, UserName, GroupName, SSHDatadir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	result, stdOut, err = ChmodSSH(sshConn, SSHDatadir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	result, stdOut, err = MkdirSSH(sshConn, SSHLog)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	result, stdOut, err = ChownSSH(sshConn, UserName, GroupName, SSHLog)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	result, stdOut, err = ChmodSSH(sshConn, SSHLog)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}

	result, stdOut, err = InitMysqlSSH(sshConn, MySQLDPath, UserName, BaseDirPath, SSHDatadir)
	if err != nil {
		log.Warnf("return code: %d: %s", result, stdOut)
	}
	log.Info("==========Initial mysqld complete==========")
	log.Info("==========Finish==========")

}
