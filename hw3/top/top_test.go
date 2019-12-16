package top

import "testing"

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
