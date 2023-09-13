package main

import (
	"fmt"
	"strconv"
)

var (
	usage = `usage:
-l line_count [file]
-b byte_count [file]
-n chunk_count [file]
`
)

type CLIOptions struct {
	ByteCount  int
	LineCount  int
	ChunkCount int
}

func NewCLIOption(byteCount int, lineCount int, chunkCount int) *CLIOptions {
	return &CLIOptions{
		ByteCount:  byteCount,
		LineCount:  lineCount,
		ChunkCount: chunkCount,
	}
}

func (opts *CLIOptions) Handle(args []string, tailArgs []string) error {
	if len(args) == 0 || len(args) != 3 || args[0] == "help" {
		fmt.Print(usage)
		return nil
	}

	split := NewSplit(tailArgs[0])
	switch args[0] {
	case "-b":
		err := split.ByByte(opts.ByteCount)
		if err != nil {
			return err
		}
	case "-l":
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if arg <= 0 {
			return fmt.Errorf("%d: illegal line count", arg)
		}
		err = split.ByLine(opts.LineCount)
		if err != nil {
			return err
		}
	case "-n":
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if arg <= 0 {
			return fmt.Errorf("%d: illegal line count", arg)
		}
		err = split.ByChunk(opts.ChunkCount)
		if err != nil {
			return err
		}
	default:
		err := split.ByLine(1000)
		if err != nil {
			return err
		}
	}
	return nil
}
