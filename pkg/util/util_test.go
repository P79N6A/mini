package util

import "testing"

func TestSaveToDisk(t *testing.T) {
	path := "../../output"
	data := "aadawdawdawd"
	url := "test"
	err := SaveToDisk(data, url, path)
	if err != nil {
		t.Errorf("SaveToDisk err, err: %s", err)
	}
}
