package trivy

import (
	"os"
	"os/exec"
)

func TrivyScan(pathJson string, pathDocker string, filename string) {
	//Code to Run Trivy In Golang
	cmdUpload := exec.Command("trivy", "config", "-f", "json", "-o", pathJson, "/resultsImage.json", filename)
	cmdUpload.Dir = pathDocker
	cmdUpload.Stdout = os.Stdout
	cmdUpload.Run()
}
