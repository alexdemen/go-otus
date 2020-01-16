package main

import "os"

func copy(src, dst string, offset, limit int) error {

	return nil
}

func readFile(path string, offset int64, limit int) (chan []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if offset > 0 {
		_, err := file.Seek(offset, 0)
		return nil, err
	}

	data := make(chan []byte)
	go func() {
		blockSize := 1024 * 1024
		if limit != -1 && limit < blockSize {
			blockSize = limit
		}
		dataBlock := make([]byte, 0, blockSize)
		for {
			file.Read(dataBlock)
			select {
			case data <- dataBlock:
			}
		}
	}()

	return data, nil
}
