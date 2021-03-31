package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

//Create a directory to test size
func init() {
	os.Mkdir("size_test_dir", 0777)
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("size_test_dir/test%d.txt", i)
		file, _ := os.Create(path)
		for j := 0; j < 100; j++ {
			file.WriteString("asdasdasasdasdasdasdasdasdasdasdasdasdasd \n")
		}
		defer file.Close()
	}
}

func TestDirSize(t *testing.T) {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	sizeTestDir := fmt.Sprintf("%s/size_test_dir", basepath)
	var expected int64 = 43000
	size, err := DirSize(sizeTestDir)
	if err != nil {
		t.Error(err)
	}

	if size != expected {
		t.Errorf("Expected %d, got %d", expected, size)
	}
}

func TestSizeInMB(t *testing.T) {
	tests := []struct {
		arg1     int64
		expected float64
	}{
		{721000000, 721},
		{892145714, 892.145714},
		{0, 0},
		{-21, 0},
		{123987458934, 123987.458934},
	}

	for _, test := range tests {
		if output := SizeInMB(test.arg1); output != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, output)
		}
	}
}

func TestSizeInGB(t *testing.T) {
	// fmt.Println(SizeInGB(892145714))
	tests := []struct {
		arg1     int64
		expected float64
	}{
		{721000000, 0.721},
		{892145714, 0.892145714},
		{0, 0},
		{-21, 0},
		{123987458934, 123.98745893399999},
	}
	for _, test := range tests {
		if output := SizeInGB(test.arg1); output != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, output)
		}
	}
}

func TestMBToGB(t *testing.T) {
	tests := []struct {
		arg1     float64
		expected float64
	}{
		{1000, 1},
		{-1234, 0},
		{2591, 2.591},
		{0, 0},
		{12345, 12.345},
	}
	for _, test := range tests {
		if output := SizeMBToGB(test.arg1); output != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, output)
		}
	}
}

func TestRoundGB(t *testing.T) {
	tests := []struct {
		arg1     float64
		expected float64
	}{
		{2.591, 2.59},
		{4.91234, 4.91},
		{0, 0},
		{-123, 0},

		{3.3123, 3.31},
	}
	for _, test := range tests {
		if output := RoundGB(test.arg1); output != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, output)
		}
	}
}

func TestSizeMsg(t *testing.T) {
	var buf bytes.Buffer
	flags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(flags)
	}()

	tests := []struct {
		arg1     float64
		expected string
	}{
		{1234, "Successfully freed up 1.23 GB."},
		{0, "Successfully freed up 0.0 MB."},
		{-123, "Successfully freed up 0 MB."},
		{681, "Successfully freed up 681.0 MB."},
		{34913, "Successfully freed up 34.91 GB."},
	}

	for _, test := range tests {
		SizeMsg(test.arg1)
		res := strings.TrimSpace(buf.String())
		if res != test.expected {
			t.Errorf("Test failed.\n Expected %s, got %s", test.expected, res)
		}
		buf.Reset()
	}
}
