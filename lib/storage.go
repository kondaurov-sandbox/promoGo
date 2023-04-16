package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	lineBufferSize = 4096
	readBufferSize = 64 * 1024
)

var offsets []int64

var file *os.File

func InitStorage(fileName string, duration time.Duration) {
	log.Printf("Initiating file '%s'", fileName)
	var err error
	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}

	for {
		file.Seek(0, io.SeekStart)
		buildOffsetIndex(file)
		time.Sleep(duration)
	}

}

func CloseStorage() {
	file.Close()
}

func GetLineById(lineNum int) (string, error) {
	if lineNum < 1 || lineNum > len(offsets) {
		return "", fmt.Errorf("Line number out of range: %d\n", lineNum)
	}
	offset := offsets[lineNum-1]
	// Seek to the offset of the line and read it using buffered I/O
	file.Seek(offset, io.SeekStart)
	lineReader := bufio.NewReaderSize(file, lineBufferSize)
	line, err := lineReader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	if line == "" {
		return "", fmt.Errorf("EOF")
	}
	return line, nil
}

func buildOffsetIndex(file *os.File) {
	var lastOffset int64 = 0
	offsets = []int64{lastOffset}
	buf := make([]byte, readBufferSize)
	log.Print("Building offsets array has been started")
	for {
		// Read a chunk of data from the file
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// Scan the chunk for newline characters and update the offset index
		for i := 0; i < n; i++ {
			lastOffset++
			if buf[i] == '\n' {
				offsets = append(offsets, lastOffset)
			}
		}

		// If we've reached the end of the file, exit the loop
		if err == io.EOF {
			break
		}
	}
	log.Print("Building offsets array has been finished")
}
