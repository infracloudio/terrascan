package template

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestIfTemplateStructIsInlineWithARMTemplate(t *testing.T) {
	data, err := readTemplate("storage-account-create.json")
	if err != nil {
		t.Error(err)
	}

	var tmp Template
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		t.Error(err)
	}
}

func readTemplate(name string) ([]byte, error) {
	const testData = "test_data"

	f, err := os.Open(filepath.Join(testData, name))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
