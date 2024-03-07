package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainWithCFlag(t *testing.T) {
	testFlag(t, "-c", "Hello, World!", 13)
	testFlag(t, "-c", "Hello, 世界", 13)
}

func TestMainWithLFlag(t *testing.T) {
	testFlag(t, "-l", "Hello, World!\nHello, World!", 2)
}

func TestMainWithWFlag(t *testing.T) {
	testFlag(t, "-w", "Hello, World!\nHello, World!", 4)
}

func TestMainWithMFlag(t *testing.T) {
	testFlag(t, "-m", "Hello, World!", 13)
	testFlag(t, "-m", "Hello, 世界", 9)
}

func TestMainWithNoFlag(t *testing.T) {
	testFlag(t, "", "Hello, World!\nThis is a test.", 28, 2, 6)
}

func testFlag(t *testing.T, flagStr string, text string, expected ...int) {
	tmpfile, err := os.CreateTemp("", "testExample")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"", flagStr, tmpfile.Name()}
	main()

	w.Close()
	os.Stdout = old

	newOut := &bytes.Buffer{}
	newOut.ReadFrom(r)
	res := newOut.String()

	expectedOutput := fmt.Sprintf("%s %s\n", strings.TrimSuffix(strings.Repeat("%d ", len(expected)), " "), tmpfile.Name())
	expectedIface := make([]interface{}, len(expected))
	for i, v := range expected {
		expectedIface[i] = v
	}
	expectedOutput = fmt.Sprintf(expectedOutput, expectedIface...)
	if res != expectedOutput {
		t.Errorf("Expected '%s', got '%s'", expectedOutput, res)
	}
}
