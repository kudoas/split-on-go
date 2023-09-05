package main

import (
	"fmt"
	"strconv"
)

type CLIOptions struct {
	ByteCount  int
	LineCount  int
	ChunkCount int
}

func (opts *CLIOptions) Handle(args []string, tailArgs []string) error {
	if len(args) == 0 || len(tailArgs) != 1 || len(args) > 0 && args[0] == "help" {
		usage := `usage:
-l line_count [file]
-b byte_count [file]
-n chunk_count [file]
`
		fmt.Print(usage)
		return nil
	}
	switch args[0] {
	case "-b":
		if len(args) != 3 {
			return fmt.Errorf("illegal option, Please check with `help` cmd.")
		}
		splitByBytes(tailArgs[0], opts.ByteCount)
	case "-l":
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if arg <= 0 {
			return fmt.Errorf("%d: illegal line count", arg)
		}
		if len(args) != 3 {
			return fmt.Errorf("illegal option, Please check with `help` cmd.")
		}
		splitByLines(tailArgs[0], opts.LineCount)
	case "-n":
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if arg <= 0 {
			return fmt.Errorf("%d: illegal line count", arg)
		}
		if len(args) != 3 {
			return fmt.Errorf("illegal option, Please check with `help` cmd.")
		}
		splitByChunks(tailArgs[0], opts.ChunkCount)
	default:
		splitByLines(tailArgs[0], 1000)
	}
	return nil
}
