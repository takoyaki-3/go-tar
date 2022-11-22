package gotar

import (
	"archive/tar"
	filetool "github.com/takoyaki-3/file-tool"
	"io"
	"os"
	"path/filepath"
)

func GetFilelist(tarPath string) (files []filetool.FileInfo, err error) {

	//
	file, err := os.Open(tarPath)
	if err != nil {
		return files, err
	}
	defer file.Close()

	// tarの展開
	tarReader := tar.NewReader(file)

	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		files = append(files, filetool.FileInfo{
			Name: tarHeader.Name,
		})
	}

	return files, nil
}

func UnTar(tarPath string, unTarPath string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// tarの展開
	tarReader := tar.NewReader(file)

	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		os.MkdirAll(filepath.Dir(unTarPath+"/"+tarHeader.Name), 0777)
		file, err := os.Create(unTarPath + "/" + tarHeader.Name)
		if err != nil {
			return err
		}
		io.Copy(file, tarReader)
	}
	return nil
}
