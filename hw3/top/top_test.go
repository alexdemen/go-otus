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
		{"Test one - success.", []string{"test", "one", "-", "success"}},
	}

	var res []string
	for _, testCase := range testCases {
		res = prepareText(testCase.input)

		if len(res) != len(testCase.output) {
			t.Errorf("Тест завершился неудачей при входном значении - \"%s\"", testCase.input)
		}
	}
}

func TestTop(t *testing.T){
	testCases := []struct{
		input string
		count int
		output []string
	}{
		{"1 1, - - - - - -  1 2 2 3 3 4", 3, []string{"one", "two", "apple"}},
	}

	for _, testCase := range testCases{
		res := Top(testCase.input, testCase.count)

		fmt.Println(res)
	}
}
