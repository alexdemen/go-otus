package main

import (
	"fmt"
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input  string
		result string
		err    error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", fmt.Errorf("")},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\45", "qwe44444", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	}

	for _, curCase := range testCases {
		if result, err := unpack(curCase.input); result != curCase.result {
			t.Errorf("The result of the method \"unpack\" with params str = \"%s\" is equal \"%s\","+
				"but not equal to expected value \"%s\"",
				curCase.input, result, curCase.result)
		} else if err == nil && curCase.err != nil {
			t.Errorf("Error expected.")
		}
	}
}

func TestRepeatRune(t *testing.T)  {
	testCases := []struct{
		count int
		symbol rune
	}{
		{2, 't'},
	}

	var runes []rune = nil

	for _, curCase := range testCases{
		runes = make([]rune, 0, curCase.count)
		repeatRune(&runes, curCase.symbol, curCase.count)

		if len(runes) != curCase.count{
			t.Errorf("The expected length value does not match. Count: %d, symbol: %c", curCase.count, curCase.symbol)
			continue
		}

		for i := 0; i < len(runes); i++{
			if runes[i] != curCase.symbol{
				t.Errorf("The expected symbol does not match. Count: %d, symbol: %c", curCase.count, curCase.symbol)
				break
			}
		}
	}
}
