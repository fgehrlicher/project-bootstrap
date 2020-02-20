package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
