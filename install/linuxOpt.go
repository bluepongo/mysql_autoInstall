package install

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/romberli/log"
)

const (
	LogFilePath = "/tmp/test.log"
)

var out bytes.Buffer
var stderr bytes.Buffer

// Execute the linux command
func ExecuteCommand(command string) (output string, err error) {

	// Initialize a logger
	fileName := LogFilePath
	_, _, err = log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		fmt.Printf("Init logger failed.\n%s", err.Error())
	}

	var stdoutBuffer bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stdoutBuffer

	err = cmd.Run()
	if err != nil {
		log.Warnf("%s: %s", err, stderr.String())
		fmt.Printf("[Warn]%s,%s: %s\n", command, err, stderr.String())
		fmt.Printf("out:%s\n", out.String())
		return stdoutBuffer.String(), err
	}
	log.Infof("%s success. %s\n", command, out.String())
	fmt.Printf("[Info]%s success. %s\n", command, out.String())
	return stdoutBuffer.String(), err
}

// Add a new group
func AddGroup(groupName string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo groupadd -g 700 %s", groupName))
}

// Add a new user and assign him to the group
func AddUser(groupName, userName string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo useradd -u 700 -g %s %s", userName, groupName))
}

// Chown command
func Chown(groupName, userName, chPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo chown %s %s %s", "-R", groupName+":"+userName, chPath))
}

// Chmod command
func Chmod(chPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo chmod %s %s %s", "-R", "775", chPath))
}

// Create a new file
func Mkdir(targetPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo mkdir -p %s", targetPath))
}

// Move a file to the toPath
func Mv(fromPath, toPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo mv -f %s %s", fromPath, toPath))
}

// Copy a file to the toPath
func Cp(fromPath, toPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo cp -rf %s %s", fromPath, toPath))
}

// Search for the file content
func Cat(targetPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo cat %s", targetPath))
}

// Establish a soft connection
func Ln(fromPath, toPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo ln -s %s %s", fromPath, toPath))
}

// Delete the target file
func Rm(targetPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo rm -f %s", targetPath))
}

// Start a service
func ServiceStart(serviceName string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo %s start", serviceName))
}

// Restart a service
func ServiceRestart(serviceName string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("sudo service %s restart", serviceName))
}

func InitMysql(mysqldPath, userName, dataDirPath, baseDirPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf(
			"sudo %s --initialize-insecure --user=%s --datadir=%s --basedir=%s",
			mysqldPath, userName, dataDirPath, baseDirPath))
}

func MultiInitMysql(mysqldPath, userName, baseDirPath, dataDirPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf(
			"runuser -l mysql -c '%s --initialize-insecure --user=%s --basedir=%s --datadir=%s'",
			mysqldPath, userName, baseDirPath, dataDirPath))
}

func MultiStartMysql(portNum string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf(
			"runuser -l mysql -c 'mysqld_multi start %s'", portNum))
}

func MultiStopMysql(portNum string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf(
			"runuser -l mysql -c 'mysqld_multi stop %s'", portNum))
}
