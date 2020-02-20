package config

type Config struct {
	JenkinsConfig JenkinsConfig `json:"jenkins"`
	GitlabConfig  GitlabConfig `json:"gitlab"`
}

type JenkinsConfig struct {
	Url string `json:"url"`
}

type GitlabConfig struct {
	Url string `json:"url"`
}
