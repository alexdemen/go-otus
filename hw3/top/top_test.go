package top

import "testing"

func TestPrepareText(t *testing.T){
	testCases := []struct{
		input string
		output []string
	}{
		{"Test one - success.", []string{"test", "one", "success"}},
	}
}
