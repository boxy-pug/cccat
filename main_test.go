package main

import (
	"os/exec"
	"testing"
)

func TestCatCloneCLI(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "./testdata/test3.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	expected := `hello
goodbye

yes man great!
`
	if string(output) != expected {
		t.Errorf("Expected %q, got %q", expected, string(output))
	}
}

func TestMultipleFiles(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "./testdata/test3.txt", "./testdata/test4.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	expected := `hello
goodbye

yes man great!
checking 

this is good
`
	if string(output) != expected {
		t.Errorf("Expected %q, got %q", expected, string(output))
	}
}

func TestNumberedLines(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-n", "./testdata/test3.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	expected := `1 hello
2 goodbye
3 
4 yes man great!
`
	if string(output) != expected {
		t.Errorf("Expected %q, got %q", expected, string(output))
	}
}
