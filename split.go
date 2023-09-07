package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Split struct {
	filePath string
}

func NewSplit(filePath string) *Split {
	return &Split{
		filePath: filePath,
	}
}

func (s *Split) ByByte(bytesPerFile int) error {
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if bytesPerFile == 0 {
		outputFile, err := os.Create("1")
		if err != nil {
			return err
		}
		_, err = io.Copy(outputFile, file)
		if err != nil {
			return err
		}
		return nil
	}

	buffer := make([]byte, bytesPerFile)
	r := bufio.NewReader(file)

	for i := 1; ; i++ {
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		outputFileName := fmt.Sprintf("%d", i)
		err = os.WriteFile(outputFileName, buffer[:n], 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Split) ByLine(linesPerFile int) error {
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	i := 1
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		if len(lines) >= linesPerFile {
			outputFileName := fmt.Sprintf("%d", i)
			outputFile, err := os.Create(outputFileName)
			if err != nil {
				return err
			}
			for _, line := range lines {
				outputFile.WriteString(line + "\n")
			}
			err = outputFile.Close()
			if err != nil {
				return err
			}
			lines = nil
			i++
		}
	}

	if len(lines) > 0 {
		outputFileName := fmt.Sprintf("%d", i)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		for _, line := range lines {
			outputFile.WriteString(line + "\n")
		}
		err = outputFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Split) ByChunk(chunksPerFile int) error {
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	chunkSize := fileInfo.Size() / int64(chunksPerFile)
	if chunkSize == 0 {
		return fmt.Errorf("can't split into more than %d files", fileInfo.Size())
	}

	buffer := make([]byte, chunkSize)
	r := bufio.NewReader(file)
	for i := 1; i <= chunksPerFile; i++ {
		outputFileName := fmt.Sprintf("%d", i)
		_, err = r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		err = os.WriteFile(outputFileName, buffer, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}
