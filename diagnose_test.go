package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

//Create directories to test
func init() {
	os.Mkdir("diagnose_test_dir", 0777)
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("diagnose_test_dir/test%d.bar", i)
		file, _ := os.Create(path)
		if err := file.Truncate(1e7); err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}
}

func TestDiagnoseDir(t *testing.T) {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	diagnoseTestDir := fmt.Sprintf("%s/diagnose_test_dir", basepath)

	rescue := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	diagnoseDir(diagnoseTestDir, "diagnose_test_dir")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescue
	str := string(out)
	expected := fmt.Sprintf("diagnose_test_dir: ----------------------- %s", green("100.0 MB"))
	res := strings.TrimSpace(str)

	if res != expected {
		t.Errorf("Test failed.\nExpected %s\nGot %s", expected, res)
	}
}

func TestAddPadding(t *testing.T) {
	tests := []struct {
		arg1     string
		expected string
	}{
		{"EFLKL+!G+", "-------------------------------"},
		{"QWERTY", "----------------------------------"},
		{"ASDELJQCNG", "------------------------------"},
		{"", "----------------------------------------"},
	}

	for _, test := range tests {
		if addPadding(test.arg1) != test.expected {
			t.Fail()
		}
	}
}
