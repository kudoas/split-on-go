package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	byteOption := flag.Int("b", 0, "-b: Create split files byte_count bytes in length.")
	lineOption := flag.Int("l", 0, "-l: Create split files line_count lines in length.")
	chunkOption := flag.Int("n", 0, "-n: Split file into chunk_count smaller files.")
	flag.Parse()
	tailArgs := flag.Args()

	cliOptions := NewCLIOption(*byteOption, *lineOption, *chunkOption)
	err := cliOptions.Handle(os.Args[1:], tailArgs)
	if err != nil {
		fmt.Println(err)
	}
}
