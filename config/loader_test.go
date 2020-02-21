package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type InvalidReader struct{}

var errInvalidReader = errors.New("im an invalid reader")

func (i InvalidReader) Read(p []byte) (n int, err error) {
	err = errInvalidReader
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
	jenkinsURL := "https://jenkins.com"
	gitlabURL := "https://gitlab.com"
	jsonTemplate := loadFixture(t, "valid_config.json")

	configJSON := fmt.Sprintf(jsonTemplate, jenkinsURL, gitlabURL)
	reader := strings.NewReader(configJSON)
	assertion := assert.New(t)

	fileLoader := Loader{}
	config, err := fileLoader.Load(reader)
	assertion.Nil(err)

	assertion.Equal(jenkinsURL, config.JenkinsConfig.Url)
	assertion.Equal(gitlabURL, config.GitlabConfig.Url)
}

func TestLoaderLoadInvalidJson(t *testing.T) {
	invalidJSON := loadFixture(t, "invalid_config.txt")
	reader := strings.NewReader(invalidJSON)
	_, err := Loader{}.Load(reader)
	assert.NotNil(t, err)
}

func TestLoaderLoadNilReader(t *testing.T) {
	_, err := Loader{}.Load(InvalidReader{})
	assert.Equal(t, errInvalidReader, err)
}
