package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestWalk(t *testing.T) {
	secondWalk := func(targetTest string) {
		need := walkDir("./")
		if need {
			t.Fatal("second walk after", targetTest, "should get false")
		}
	}
	fatalErr := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	// init some files
	err := os.MkdirAll("./walk_test", os.ModePerm)
	fatalErr(err)
	time.Sleep(time.Second * 2)
	defer os.RemoveAll("./walk_test")

	// start test
	walkLog = true

	// init walk
	initWalk("./")
	secondWalk("init walk")

	// new file
	err = ioutil.WriteFile("./walk_test/file", []byte{}, 0777)
	fatalErr(err)
	time.Sleep(time.Second)
	need := walkDir("./")
	if !need {
		t.Fatal("new file should get true")
	}
	secondWalk("new")

	// modify file
	err = ioutil.WriteFile("./walk_test/file", []byte{123}, 0777)
	fatalErr(err)
	time.Sleep(time.Second)
	need = walkDir("./")
	if !need {
		t.Fatal("modify file should get true")
	}
	secondWalk("modify")

	// delete file
	err = os.Remove("./walk_test/file")
	fatalErr(err)
	time.Sleep(time.Second)
	need = walkDir("./")
	if !need {
		t.Fatal("del file should get true")
	}
	secondWalk("del")
}
