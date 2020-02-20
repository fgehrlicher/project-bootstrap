package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type InvalidReader struct {}
var invalidReaderError = errors.New("im an invalid reader")

func (i InvalidReader) Read(p []byte) (n int, err error) {
	err = invalidReaderError
	return
}

func loadFixture(t *testing.T, name string) string {
	content, err := ioutil.ReadFile("./testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}

	return string(content)
}

func TestLoaderLoad(t *testing.T) {
	jenkinsUrl := "https://jenkins.com"
	gitlabUrl := "https://gitlab.com"
	jsonTemplate := loadFixture(t, "valid_config.json")

	configJson := fmt.Sprintf(jsonTemplate, jenkinsUrl, gitlabUrl)
	reader := strings.NewReader(configJson)
	assertion := assert.New(t)

	fileLoader := Loader{}
	config, err := fileLoader.Load(reader)
	assertion.Nil(err)

	assertion.Equal(jenkinsUrl, config.JenkinsConfig.Url)
	assertion.Equal(gitlabUrl, config.GitlabConfig.Url)
}

func TestLoaderLoadInvalidJson(t *testing.T) {
	invalidJson := loadFixture(t, "invalid_config.txt")
	reader := strings.NewReader(invalidJson)
	_, err := Loader{}.Load(reader)
	assert.NotNil(t, err)
}

func TestLoaderLoadNilReader(t *testing.T) {
	_, err := Loader{}.Load(InvalidReader{})
	assert.Equal(t, invalidReaderError, err)
}
