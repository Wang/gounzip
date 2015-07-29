package unzip

import (
	"os"
	"testing"
)

func TestUnzip(t *testing.T) {
	if err := Do("", ""); err == nil {
		t.Error(err.Error())
	}

	if err := Do("./test.zip", ""); err != nil {
		t.Error(err.Error())
	} else if _, err := os.Stat("./file1"); err != nil {
		t.Error(err.Error())
	}
	os.Remove("./file1")

	if err := Do("./test.zip", "test"); err != nil {
		t.Error(err.Error())
	} else if _, err := os.Stat("./test/file1"); err != nil {
		t.Error(err.Error())
	}
	os.RemoveAll("./test")
}
