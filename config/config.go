package config

// Config is the main configuration
type Config struct {
	JenkinsConfig JenkinsConfig `yaml:"jenkins"`
	GitlabConfig  GitlabConfig  `yaml:"gitlab"`
}

// JenkinsConfig is the jenkins specific config
type JenkinsConfig struct {
	URL string `json:"url"`
}

// GitlabConfig is the gitlab specific config
type GitlabConfig struct {
	URL string `json:"url"`
}
