package enviroment

import (
	"io/ioutil"
	"os"
)

func ReadDir(dir string) (map[string]string, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	res := make(map[string]string)

	for _, file := range files {
		val, err := readFile(dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			return nil, err
		}

		res[file.Name()] = val
	}

	return res, nil
}

func readFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
