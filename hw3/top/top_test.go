package top

import (
	"fmt"
	"testing"
)

func TestPrepareText(t *testing.T) {
	testCases := []struct {
		input  string
		output []string
	}{
		{"Test one - success.", []string{"test", "one", "success"}},
	}

	var res []string
	for _, testCase := range testCases {
		res = prepareText(testCase.input)

		if len(res) != len(testCase.output) {
			t.Errorf("Тест завершился неудачей при входном значении - \"%s\"", testCase.input)
		}
	}
}

func TestTop(t *testing.T) {
	testCases := []struct {
		input  string
		count  int
		output []string
	}{
		{"one two two three three three", 2, []string{"three", "two"}},
		{"one two two. - three - three, - three", 2, []string{"three", "two"}},
	}

	for _, testCase := range testCases {
		res := Top(testCase.input, testCase.count)

		if len(res) != len(testCase.output) {
			t.Errorf("Тест завершился неудачей при входном значении - \"%s\"", testCase.input)
			continue
		}

		for index, val := range res {
			if val != testCase.output[index] {
				t.Errorf("Тест завершился неудачей при входном значении - \"%s\"", testCase.input)
			}
		}

		fmt.Println(res)
	}
}
