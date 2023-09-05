package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSplitByBytes(t *testing.T) {
	cases := []struct {
		content       string
		bytesPerFile  int
		expectedFiles []string
		expectedError bool
	}{
		{
			content:       "larger than bytesPerFile",
			bytesPerFile:  10,
			expectedFiles: []string{"1", "2", "3"},
			expectedError: false,
		},
		{
			content:       "smaller than bytesPerFile",
			bytesPerFile:  100000,
			expectedFiles: []string{"1"},
			expectedError: false,
		},
		{
			content:       "0 byte",
			bytesPerFile:  0,
			expectedFiles: []string{"1"},
			expectedError: false,
		},
	}

	for _, c := range cases {
		t.Run(c.content, func(t *testing.T) {
			inputPath := "test_input.txt"
			err := createTestFile(t, inputPath, c.content)
			if err != nil {
				t.Errorf("Failed to create test input file: %v", err)
			}
			defer os.Remove(inputPath)

			err = splitByBytes(inputPath, c.bytesPerFile)
			if err != nil && !c.expectedError {
				t.Errorf("splitByBytes returned an unexpected error: %v", err)
			} else if err == nil && c.expectedError {
				t.Errorf("Expected error, but got none")
			}
			for _, expectedFile := range c.expectedFiles {
				_, err := os.Stat(expectedFile)
				if err != nil {
					t.Errorf("File %s was not created as expected: %v", expectedFile, err)
				}
			}
			verifyFileSize(t, c.bytesPerFile, c.expectedFiles)
			cleanTestFile(t, c.expectedFiles)
		})
	}
}

func TestSplitByLines(t *testing.T) {
	cases := []struct {
		content       string
		linesPerFile  int
		expectedFiles []string
		expectedError bool
	}{
		{
			content:       "longer than linesPerFile \n rows",
			linesPerFile:  1,
			expectedFiles: []string{"1", "2"},
			expectedError: false,
		},
		{
			content:       "shorter than linesPerFile \n rows",
			linesPerFile:  3,
			expectedFiles: []string{"1"},
			expectedError: false,
		},
	}

	for _, c := range cases {
		t.Run(c.content, func(t *testing.T) {
			inputPath := "test_input.txt"
			err := createTestFile(t, inputPath, c.content)
			if err != nil {
				t.Errorf("Failed to create test input file: %v", err)
			}
			defer os.Remove(inputPath)

			err = splitByLines(inputPath, c.linesPerFile)
			if err != nil && !c.expectedError {
				t.Errorf("splitByLines returned an unexpected error: %v", err)
			} else if err == nil && c.expectedError {
				t.Errorf("Expected error, but got none")
			}
			for _, expectedFile := range c.expectedFiles {
				_, err := os.Stat(expectedFile)
				if err != nil {
					t.Errorf("File %s was not created as expected: %v", expectedFile, err)
				}
			}
			verifyFileLine(t, c.linesPerFile, c.expectedFiles)
			cleanTestFile(t, c.expectedFiles)
		})
	}
}

func TestSplitByChunks(t *testing.T) {
	cases := []struct {
		content       string
		chunksPerFile int
		expectedFiles []string
		expectedError bool
	}{
		{
			content:       "split by chunk",
			chunksPerFile: 3,
			expectedFiles: []string{"1", "2", "3"},
			expectedError: false,
		},
		{
			content:       "...",
			chunksPerFile: 100,
			expectedFiles: []string{},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.content, func(t *testing.T) {
			inputPath := "test_input.txt"
			err := createTestFile(t, inputPath, c.content)
			if err != nil {
				t.Errorf("Failed to create test input file: %v", err)
			}
			defer os.Remove(inputPath)

			err = splitByChunks(inputPath, c.chunksPerFile)
			if err != nil && !c.expectedError {
				t.Errorf("splitByChunks returned an unexpected error: %v", err)
			} else if err == nil && c.expectedError {
				t.Errorf("Expected error, but got none")
			}
			for _, expectedFile := range c.expectedFiles {
				_, err := os.Stat(expectedFile)
				if err != nil {
					t.Errorf("File %s was not created as expected: %v", expectedFile, err)
				}
			}
			verifyFileCount(t, c.expectedFiles)
			cleanTestFile(t, c.expectedFiles)
		})
	}
}

func createTestFile(t *testing.T, filePath, content string) error {
	t.Helper()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func verifyFileSize(t *testing.T, expectedFileSize int, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file %s: %v", fileName, err)
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			t.Error(err)
		}
		if expectedFileSize != 0 && int(info.Size()) > expectedFileSize {
			t.Errorf("File %s size mismatch, Expected: %d, but got: %d", fileName, expectedFileSize, info.Size())
		}
	}
}

func verifyFileLine(t *testing.T, expectedLineCount int, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file %s: %v", fileName, err)
		}
		scanner := bufio.NewScanner(file)
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}
		if err := scanner.Err(); err != nil {
			t.Errorf("Error reading file %s: %v", fileName, err)
		}
		if int(lineCount) > expectedLineCount {
			t.Errorf("File %s line mismatch, Expected: %d, but got: %d", fileName, expectedLineCount, lineCount)
		}
	}
}

func verifyFileCount(t *testing.T, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("File %s count mismatch", fileName)
		}
		defer file.Close()
	}
}

func cleanTestFile(t *testing.T, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		if err := os.Remove(fileName); err != nil {
			t.Error(err)
		}
	}
}
