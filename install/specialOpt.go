package install

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/romberli/log"
)

// UnTarGz a .tar.gz file.
func UnTarGz(srcFilePath string, destDirPath string) {
	// Create destination directory
	Mkdir(destDirPath)

	// Initialize a logger
	fileName := LogFilePath
	_, _, err := log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		fmt.Printf("Init logger failed.\n%s", err.Error())
	}

	log.Info("UnTarGzing file ...")
	fr, err := os.Open(srcFilePath)
	if err != nil {
		log.Warn("The src-file is not exits.")
		return
	}
	defer fr.Close()

	// Gzip reader
	gr, err := gzip.NewReader(fr)

	// Tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			// Write data to file
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				log.Warn("Cannot Create destDirpath.")
				return
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				log.Warn("Cannot Copy.")
				return
			}
		}
	}
	fmt.Printf("[Info]UnTarGzing file '%s' successfully!\n", srcFilePath)
	log.Infof("UnTarGzing file '%s' successfully!", srcFilePath)
}

// Initialize mysqld.
func InitMysqld(mySqld string, userName string, dataDirPath string, baseDirPath string) {
	// Initialize a logger
	fileName := LogFilePath
	_, _, err := log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		fmt.Printf("Init logger failed.\n%s", err.Error())
	}

	cmd := exec.Command(
		mySqld, "--initialize", "--user="+userName, "--datadir="+dataDirPath, "--basedir="+baseDirPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Execte the command
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		log.Warn("Initialize failed.")
		return
	}

	fmt.Printf("==========...Initializing...==========\n%s\nInitialize successfully!\n", out.String())

	log.Info("Initialize successfully!")
}

// Execute the linux command
func Mv(fromPath, toPath string) (output string, err error) {

	// Initialize a logger
	fileName := LogFilePath
	_, _, err = log.InitLoggerWithDefaultConfig(fileName)
	if err != nil {
		fmt.Printf("Init logger failed.\n%s", err.Error())
	}

	var stdoutBuffer bytes.Buffer

	cmd := exec.Command(fmt.Sprintf("mv %s %s", fromPath, toPath))
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stdoutBuffer

	err = cmd.Run()
	if err != nil {
		log.Warnf("%s: %s", err, stderr.String())
		fmt.Printf("%s: %s\n", err, stderr.String())
		return stdoutBuffer.String(), err
	}
	log.Infof("mv success. %s\n", out.String())
	fmt.Printf("[Info]mv success. %s\n", out.String())
	return stdoutBuffer.String(), err
}