package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// write tests for main.go
//
// Path: main_test.go
// Compare this snippet from old_main_test.go:

// // TestMain tests the main function.
func TestMain(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if !strings.Contains(string(out), "Enter the provider's pubkey") {
		t.Errorf("Unexpected output: %s", out)
	}
}
