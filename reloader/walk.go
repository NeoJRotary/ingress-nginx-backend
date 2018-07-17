package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type walkFile struct {
	modtime time.Time
	checked bool
}

var walkMap = map[string]*walkFile{}
var hasUpdates = false
var walkLog = false

func initWalk(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		walkMap[path] = &walkFile{modtime: info.ModTime()}
		return nil
	})
}

func walkDir(dir string) bool {
	// reset
	hasUpdates = false
	for _, wf := range walkMap {
		wf.checked = false
	}

	// walk
	filepath.Walk(dir, walkFn)

	// check file deletion
	for k, wf := range walkMap {
		if !wf.checked {
			hasUpdates = true
			delete(walkMap, k)
			if walkLog {
				fmt.Println("Walk find del", k)
			}
		}
	}

	return hasUpdates
}

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(path, err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	wf, ok := walkMap[path]
	if ok {
		if !wf.modtime.Equal(info.ModTime()) {
			hasUpdates = true
			wf.modtime = info.ModTime()
			if walkLog {
				fmt.Println("Walk find mod", path)
			}
		}
	} else {
		hasUpdates = true
		wf = &walkFile{modtime: info.ModTime()}
		walkMap[path] = wf
		if walkLog {
			fmt.Println("Walk find new", path)
		}
	}
	wf.checked = true

	return nil
}
