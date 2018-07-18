package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/NeoJRotary/exec-go"
)

var (
	configmapFolder  = "/etc/config"
	configmapScanDur = 60
)

func updateENVs() {
	envStr := os.Getenv("CONFIGMAP_FOLDER")
	if envStr != "" {
		configmapFolder = envStr
	}

	envStr = os.Getenv("CONFIGMAP_SCAN_DUR")
	if envStr != "" {
		i, err := strconv.Atoi(envStr)
		if err != nil {
			panic(err)
		}
		configmapScanDur = i
	}
}

func main() {
	updateENVs()
	initWalk(configmapFolder)
	fmt.Println("Start Watching on ConfigMap Folder : ", configmapFolder)

	for {
		needUpdate := walkDir(configmapFolder)

		fmt.Println("ConfigMap Reload Checked, Need Update : ", needUpdate)

		if needUpdate {
			reload()
		}
		time.Sleep(time.Second * time.Duration(configmapScanDur))

		updateENVs()
	}
}

func reload() {
	fmt.Println("Reloader exec : sh /sync_files.sh")

	out, err := exec.NewCmd("", "sh", "/sync_files.sh").Run()
	if err != nil {
		fmt.Println(out + err.Error())
	} else {
		fmt.Println(out)
	}

	fmt.Println("Reloader exec : nginx -s reload")

	out, err = exec.NewCmd("", "nginx", "-s", "reload").Run()
	if err != nil {
		fmt.Println(out + err.Error())
	} else {
		fmt.Println(out)
	}
}
