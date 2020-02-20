package config

import (
	"errors"
	"fmt"
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

func TestLoaderLoad(t *testing.T) {
	jenkinsUrl := "https://jenkins.com"
	gitlabUrl := "https://gitlab.com"
	jsonTemplate := `
{
	"jenkins": {
		"url": "%s"
	},
	"gitlab": {
		"url": "%s"
	}
}
`
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
	invalidJson := `
{
		"url": "%s"
	},
	gitlab": {
		"url": "%s"
`
	reader := strings.NewReader(invalidJson)
	_, err := Loader{}.Load(reader)
	assert.NotNil(t, err)
}

func TestLoaderLoadNilReader(t *testing.T) {
	_, err := Loader{}.Load(InvalidReader{})
	assert.Equal(t, invalidReaderError, err)
}
