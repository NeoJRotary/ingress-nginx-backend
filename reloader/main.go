package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
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
	out, err := exec.Command("sh", "/sync_files.sh").Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}

	fmt.Println("Reloader exec : nginx -s reload")
	out, err = exec.Command("nginx", "-s", "reload").Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}
