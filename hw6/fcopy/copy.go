package main

import "os"

func copy(src, dst string, offset, limit int) error {

	return nil
}

func readFile(path string, offset int64, limit int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if offset > 0 {
		_, err := file.Seek(offset, 0)
		return err
	}
	limit := 1024 * 1024
	for {
		file.Read()
	}
	return nil
}
