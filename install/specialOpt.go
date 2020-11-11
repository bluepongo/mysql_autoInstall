package install

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
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
