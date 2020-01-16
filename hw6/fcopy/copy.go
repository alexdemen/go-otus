package main

import (
	"io"
	"os"
)

func copyFile(src, dst string, offset, limit int) error {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0775)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(src)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	blockSize := 1024 * 1024

	var readerCount int
	buffer := make([]byte, 0, blockSize)

	for isEnd := false; !isEnd && (readerCount < limit || limit == -1); {

		if limit != -1 && readerCount+blockSize > limit {
			blockSize = limit - readerCount
			buffer = make([]byte, 0, blockSize)
		}

		readSize, err := srcFile.ReadAt(buffer, int64(offset))
		if err != nil {
			if err == io.EOF {
				isEnd = true
			} else {
				return err
			}
		}

		_, err = dstFile.Write(buffer)
		if err != nil {
			return err
		}

		buffer = buffer[:0]
		readerCount += readSize
	}

	return nil
}
