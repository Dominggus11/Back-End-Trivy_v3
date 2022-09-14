package trivy

import (
	"os"
	"os/exec"
	"strconv"
)

func MkdirUpload() {
	folder := 1
	pathFolder := "FileDocker/FileUpload/"
	for i := 1; i < 10; i++ {
		temp := strconv.Itoa(folder)
		_, err := os.Stat(pathFolder + temp)
		if err == nil {
			folder = folder + 1
		} else {
			cmdUpload := exec.Command("mkdir", pathFolder+temp)
			cmdUpload.Stdout = os.Stdout
			cmdUpload.Run()
			break
		}
	}
}

func MkdirWrite() {
	folder := 1
	pathFolder := "FileDocker/FileWrite/"
	for i := 1; i < 10; i++ {
		temp := strconv.Itoa(folder)
		_, err := os.Stat(pathFolder + temp)
		if err == nil {
			folder = folder + 1
		} else {
			cmdUpload := exec.Command("mkdir", pathFolder+temp)
			cmdUpload.Stdout = os.Stdout
			cmdUpload.Run()
			break
		}
	}
}
