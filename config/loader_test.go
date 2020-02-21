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
	ymlTemplate := loadFixture(t, "valid_config.yml")

	ymlConfig := fmt.Sprintf(ymlTemplate, jenkinsURL, gitlabURL)
	reader := strings.NewReader(ymlConfig)
	assertion := assert.New(t)

	fileLoader := Loader{}
	config, err := fileLoader.Load(reader)
	assertion.Nil(err)

	assertion.Equal(jenkinsURL, config.JenkinsConfig.URL)
	assertion.Equal(gitlabURL, config.GitlabConfig.URL)
}

func TestLoaderLoadInvalidJson(t *testing.T) {
	invalidYML := loadFixture(t, "invalid_config.txt")
	reader := strings.NewReader(invalidYML)
	_, err := Loader{}.Load(reader)
	assert.NotNil(t, err)
}

func TestLoaderLoadNilReader(t *testing.T) {
	_, err := Loader{}.Load(InvalidReader{})
	assert.Equal(t, errInvalidReader, err)
}
