package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

//Create directories to test
func init() {
	os.Mkdir("clean_test_dir", 0777)
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("clean_test_dir/test%d.bar", i)
		file, _ := os.Create(path)
		if err := file.Truncate(1e7); err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}
}

func TestCleanDirContents(t *testing.T) {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	cleanTestDir := fmt.Sprintf("%s/clean_test_dir", basepath)

	CleanDirContents(cleanTestDir)
	size, err := DirSize(cleanTestDir)
	if err != nil {
		t.Error(err)
	}
	if size != 0 {
		t.Fail()
	}
}
