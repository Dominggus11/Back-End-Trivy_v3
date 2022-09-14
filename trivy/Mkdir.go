package trivy

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func MkdirDocker() {
	folder := 1
	pathFolder := "FileDocker/"
	for i := 1; i < 10; i++ {
		temp := strconv.Itoa(folder)
		_, err := os.Stat(pathFolder + temp)
		if err == nil {
			folder = folder + 1
		} else {
			fmt.Println("Tidak Ada")
			cmdUpload := exec.Command("mkdir", pathFolder+temp)
			cmdUpload.Stdout = os.Stdout
			cmdUpload.Run()
			break
		}
	}
}
