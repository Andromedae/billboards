package display 

import (
	"reflect"
	"testing"
)

func TestBestFontSize(t *testing.T) {
	expect(BestFontSize(20, 6, "hacker cup"), 3, t)
}

func expect(expected, actual interface{}, t *testing.T) bool {
	t.Log("Expecting: ", expected, " Got: ", actual)
	result := reflect.DeepEqual(expected, actual)
	if !result {
		t.Fail()
	}
	return result
}