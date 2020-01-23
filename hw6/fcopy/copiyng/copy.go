package copiyng

import (
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

func Copy(src, dst string, offset, limit int) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	pbCount, err := getPbCount(limit, src, offset)
	if err != nil {
		return err
	}

	bar := pb.StartNew(pbCount)
	defer bar.Finish()

	err = copyingProcess(limit, srcFile, offset, dstFile, bar)
	if err != nil {
		return err
	}

	return nil
}

func copyingProcess(limit int, srcFile *os.File, offset int, dstFile *os.File, bar *pb.ProgressBar) error {
	blockSize := 1024 * 1024
	var readerCount int
	buffer := make([]byte, blockSize, blockSize)

	for isEnd := false; !isEnd && (readerCount < limit || limit == 0); {
		if limit != 0 && readerCount+blockSize > limit {
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

		_, err = dstFile.Write(buffer[:readSize])
		if err != nil {
			return err
		}

		offset += readSize
		readerCount += readSize

		bar.SetCurrent(int64(readerCount))
	}
	return nil
}

func getPbCount(limit int, src string, offset int) (int, error) {
	var pbCount int

	if limit != 0 {
		pbCount = limit
	} else {
		fi, err := os.Stat(src)
		if err != nil {
			return 0, err
		}
		pbCount = int(fi.Size()) - offset
	}
	return pbCount, nil
}
