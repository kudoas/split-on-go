package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	linePtr := flag.Int("l", 0, "-l: Create split files line_count lines in length.")
	chunkPtr := flag.Int("n", 0, "-n: Split file into chunk_count smaller files.")
	bytePtr := flag.Int("b", 0, "-b: Create split files byte_count bytes in length.")
	flag.Parse()

	cliOptions := &CLIOptions{ByteCount: *bytePtr, LineCount: *linePtr, ChunkCount: *chunkPtr}
	tailArgs := flag.Args()
	err := cliOptions.Handle(os.Args[1:], tailArgs)
	if err != nil {
		fmt.Println(err)
	}
}
