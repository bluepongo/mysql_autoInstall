package install

import (
	"fmt"
	"github.com/romberli/go-util/linux"
)

// Establish the ssh connection
func EstablishSSHConnect(hostIp string, portNum int, userName, userPass string) (sshConn *linux.MySSHConn, err error) {
	// Establish ssh connection
	sshConn, err = linux.NewMySSHConn(hostIp, portNum, userName, userPass)
	if err != nil {
		return nil, err
	}
	return sshConn, nil
}

// Create the user and group by SSH
func AddUserGroupSSH(sshConn *linux.MySSHConn, userName, groupName string) (result int, stdOut string, err error) {
	result, stdOut, err = sshConn.ExecuteCommand(fmt.Sprintf("sudo groupadd -g 700 %s", userName))
	if err != nil {
		return result, stdOut, err
	}
	result, stdOut, err = sshConn.ExecuteCommand(fmt.Sprintf("sudo useradd -u 700 -g %s %s", userName, groupName))
	if err != nil {
		return result, stdOut, err
	}
	return result, stdOut, err
}

// UnTarGz the file
func UnTarLocal(sourceFolder, tarname, filename, targetname string) (stdErr string, err error) {
	stdErr, err = UnTarGz(sourceFolder+targetname, sourceFolder)
	if err != nil {
		return stdErr, err
	}
	_, stdErr, err = ExecuteCommand(fmt.Sprintf("sudo mv %s %s", sourceFolder+filename, sourceFolder+targetname))
	if err != nil {
		return stdErr, err
	}
	return stdErr, err
}

// Copy the mysql Folder to remote
func CopyMysqlToRemote(sshConn *linux.MySSHConn, sourceFolder, targetPath string) (err error) {
	err = sshConn.CopyToRemote(sourceFolder+MySQL, targetPath)
	return err
}

// Create folder
func MkdirSSH(sshConn *linux.MySSHConn, targetPath string) (result int, stdOut string, err error) {
	result, stdOut, err = sshConn.ExecuteCommand(fmt.Sprintf("sudo mkdir -p %s", targetPath))
	if err != nil {
		return result, stdOut, err
	}
	return result, stdOut, err
}

// CHown
func ChownSSH(sshConn *linux.MySSHConn, groupName, userName, chPath string) (result int, stdOut string, err error) {
	result, stdOut, err = sshConn.ExecuteCommand(fmt.Sprintf("sudo chown -R %s %s", groupName+":"+userName, chPath))
	if err != nil {
		return result, stdOut, err
	}
	return result, stdOut, err
}
func ChmodSSH(sshConn *linux.MySSHConn, chPath string) (result int, stdOut string, err error) {
	result, stdOut, err = sshConn.ExecuteCommand(fmt.Sprintf("sudo chmod -R 755 %s", chPath))
	if err != nil {
		return result, stdOut, err
	}
	return result, stdOut, err
}

func InitMysqlSSH(sshConn *linux.MySSHConn, mysqldPath, userName, baseDirPath, dataDirPath string) (result int, stdOut string, err error) {
	result, stdOut, err = sshConn.ExecuteCommand(fmt.Sprintf("runuser -l mysql -c '%s --initialize-insecure --user=%s --basedir=%s --datadir=%s'",
		mysqldPath, userName, baseDirPath, dataDirPath))
	if err != nil {
		return result, stdOut, err
	}
	return result, stdOut, err
}
