package copiyng

import (
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	srcFileName := "test"
	dstFileName := "dstTest"
	srcFile, err := os.Create(srcFileName)
	srcFile.WriteString("Тестовый файл.")
	if err != nil {
		t.Fatal("Не удалось создать файл для теста.")
	}
	srcFile.Close()
	defer func() {
		os.Remove(srcFileName)
		if _, err := os.Stat(dstFileName); !os.IsNotExist(err) {
			os.Remove(dstFileName)
		}
	}()

	err = Copy(srcFileName, dstFileName, 0, 0)
	if err != nil {
		t.Errorf("Ошибка копирования файла: %s", err.Error())
	}

}
