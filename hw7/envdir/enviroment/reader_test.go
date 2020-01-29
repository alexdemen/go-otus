package enviroment

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	dirName := "test"
	os.Mkdir(dirName, 774)
	defer os.Remove(dirName)
	ioutil.WriteFile(dirName+string(os.PathSeparator)+"ABC", []byte("true"), 744)
	defer os.Remove(dirName + string(os.PathSeparator) + "ABC")

	env, err := ReadDir(dirName)
	if err != nil {
		t.Error("Ошибка чтения директории.")
	}

	if val, ok := env["ABC"]; ok {
		if val != "true" {
			t.Error("Ошибка получения переменной окружения.")
		}
	} else {
		t.Error("Ошибка получения переменной окружения.")
	}

}
