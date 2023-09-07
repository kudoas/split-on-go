package main

import (
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	cases := []struct {
		args          []string
		tailArgs      []string
		opts          CLIOptions
		expectedError bool
	}{
		{
			args:          []string{},
			tailArgs:      []string{},
			opts:          CLIOptions{},
			expectedError: false,
		},
		{
			args:          []string{"help"},
			tailArgs:      []string{"help"},
			opts:          CLIOptions{},
			expectedError: false,
		},
		{
			args:          []string{"test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{},
			expectedError: false,
		},
		{
			args:          []string{"test/test.txt", "sub"},
			tailArgs:      []string{"test/test.txt", "sub"},
			opts:          CLIOptions{},
			expectedError: false,
		},
		{
			args:          []string{"-b", "test/test.txt", "sub"},
			tailArgs:      []string{"test/test.txt", "sub"},
			opts:          CLIOptions{},
			expectedError: false,
		},
		{
			args:          []string{"-b", "10", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{ByteCount: 10},
			expectedError: false,
		},
		{
			args:          []string{"-b", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{},
			expectedError: true,
		},
		{
			args:          []string{"-n", "10", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{ChunkCount: 10},
			expectedError: false,
		},
		{
			args:          []string{"-n", "0", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{ChunkCount: 0},
			expectedError: true,
		},
		{
			args:          []string{"-l", "10", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{LineCount: 10},
			expectedError: false,
		},
		{
			args:          []string{"-l", "10", "-n", "10", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{LineCount: 10, ChunkCount: 10},
			expectedError: true,
		},
		{
			args:          []string{"-l", "0", "test/test.txt"},
			tailArgs:      []string{"test/test.txt"},
			opts:          CLIOptions{LineCount: 0},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.args, "_"), func(t *testing.T) {
			err := c.opts.Handle(c.args, c.tailArgs)
			if err != nil && !c.expectedError {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && c.expectedError {
				t.Errorf("Expected error, but got none")
			}
		})
	}
}
