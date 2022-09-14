package main

import "trivy_v3/trivy"

func main() {
	// router.Router()
	trivy.MkdirUploadFile()
	trivy.MkdirUploadJson()
	trivy.MkdirWriteFile()
	trivy.MkdirWriteJson()
}
