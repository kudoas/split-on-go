package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	bytePtr := flag.Int("b", 0, "-b: Create split files byte_count bytes in length.")
	linePtr := flag.Int("l", 0, "-l: Create split files line_count lines in length.")
	chunkPtr := flag.Int("n", 0, "-n: Split file into chunk_count smaller files.")
	flag.Parse()
	tailArgs := flag.Args()

	cliOptions := NewCLIOption(*bytePtr, *linePtr, *chunkPtr)
	err := cliOptions.Handle(os.Args[1:], tailArgs)
	if err != nil {
		fmt.Println(err)
	}
}
