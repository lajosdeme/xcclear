package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/Masterminds/semver"
)

func TestMain(m *testing.M) {
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func shutdown() {
	err := os.RemoveAll("size_test_dir")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("diagnose_test_dir")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("clean_test_dir")
	if err != nil {
		fmt.Println(err)
	}
}

func TestCountOf(t *testing.T) {

	tests := []struct {
		arg1     bool
		arg2     []bool
		expected int
	}{
		{true, []bool{true, true, false, false, true, false, true, false, false, false, true, false, false}, 5},
		{false, []bool{true, true, true, true, true, true, true, true, true, true, true, true, false}, 1},
		{false, []bool{false, true, false, true, true, false, false, false, true, true, true, true, false}, 6},
		{true, []bool{false, false, false, true, true, false, false, false, true, false, false, true, false}, 4},
	}

	for _, test := range tests {
		if output := CountOf(test.arg1, test.arg2); output != test.expected {
			t.Fail()
		}
	}
}

func TestLatestVer(t *testing.T) {

	tests := []struct {
		arg1     []string
		expected *semver.Version
	}{
		{[]string{"14.2", "13.1", "10.9.1", "8.4.5", "3", "7", "14.1.9", "2.5", "abcde", ""}, latestVerTestHelper("14.2.0")},
		{[]string{"21.2", "3.4.5", "!!!", ":)", "9.13", "1.0.1", "25.4.3", "14.9.8", "23.4"}, latestVerTestHelper("25.4.3")},
		{[]string{"12.2", "0.1.5", "q", "_DS", "2.4", "11.1.1", "4.5.9", "1", "13", " ", "7"}, latestVerTestHelper("13.0.0")},
		{[]string{"9", "7.5", "áöüóő", "8.12", "24.3.0", "85", "45.1", "23.9", "123.23.1", "3.4"}, latestVerTestHelper("123.23.1")},
	}

	for _, test := range tests {
		if output := latestVer(test.arg1); output.Compare(test.expected) != 0 {
			t.Fail()
		}
	}
}

func latestVerTestHelper(str string) *semver.Version {
	ver, _ := semver.NewVersion(str)
	return ver
}
