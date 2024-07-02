package main

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type OpenFileTest struct {
	name       string
	pathToTest string
	expectErr  bool
	errText    string
}

func TestOpenJpeg(t *testing.T) {
	var testDir = t.TempDir()
	currDir, _ := filepath.Abs("./")

	var tests = []OpenFileTest{
		{"FileDoesNotExist", fmt.Sprintf("%s/%s", testDir, "fake.jpeg"), true, "no such file or directory"},
		{"FileDoesNotExist", fmt.Sprintf("%s/%s", "fake/directory", "fake.jpeg"), true, "no such file or directory"},
		{"FileExists", fmt.Sprintf("%s/%s/%s", currDir, "images", "test_image.jpg"), false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := openJpeg(tt.pathToTest)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}
		})
	}
}

func TestWriteJpeg(t *testing.T) {
	tmpDir := t.TempDir()
	samplePixels := createGray2DArray()
	testFile := fmt.Sprintf("%v/%v", tmpDir, "test.jpg")
	err := writeJpeg(samplePixels, testFile)

	if err != nil {
		t.Errorf("writing jpg returned an error: %v", err)
	}

	if _, fileErr := os.Stat(testFile); errors.Is(fileErr, os.ErrNotExist) {
		t.Errorf("file does not exist, and it should")
	}
}

func TestWriteJpegBad(t *testing.T) {
	testFile := fmt.Sprintf("%v/%v", "abc", "test.jpg")
	err := writeJpeg([][]color.Color{}, testFile)

	if err == nil {
		t.Errorf("writing jpg should have returned an error")
	}
}

func TestWriteJpegBad2(t *testing.T) {
	tmpDir := t.TempDir()
	samplePixels := createGray2DArray()
	testFile := fmt.Sprintf("%v/%v", tmpDir, "")
	err := writeJpeg(samplePixels, testFile)

	if err == nil {
		t.Errorf("writing jpg should have returned an error")
	}
}

func TestWritePng(t *testing.T) {
	tmpDir := t.TempDir()
	samplePixels := createGray2DArray()
	testFile := fmt.Sprintf("%v/%v", tmpDir, "test.png")
	err := writePng(samplePixels, testFile)

	if err != nil {
		t.Errorf("writing png returned an error: %v", err)
	}

	if _, fileErr := os.Stat(testFile); errors.Is(fileErr, os.ErrNotExist) {
		t.Errorf("file does not exist, and it should")
	}
}

func TestWritePngBad(t *testing.T) {
	testFile := fmt.Sprintf("%v/%v", "abc", "test.png")
	err := writePng([][]color.Color{}, testFile)

	if err == nil {
		t.Errorf("writing png should have returned an error")
	}
}

func TestWritePngBad2(t *testing.T) {
	tmpDir := t.TempDir()
	samplePixels := createGray2DArray()
	testFile := fmt.Sprintf("%v/%v", tmpDir, "")
	err := writePng(samplePixels, testFile)

	if err == nil {
		t.Errorf("writing png should have returned an error")
	}
}
