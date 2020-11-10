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
		log.Warnf("%s: %s\n", err, stderr.String())
		fmt.Printf("%s: %s\n", err, stderr.String())
	}
	return stdoutBuffer.String(), err
}

func Mv(fromPath, toPath string) (output string, err error) {
	return ExecuteCommand(
		fmt.Sprintf("mv %s %s", fromPath, toPath))
}

