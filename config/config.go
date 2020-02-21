package config

// Config is the main configuration that comes from loading
// a json file.
type Config struct {
	JenkinsConfig JenkinsConfig `json:"jenkins"`
	GitlabConfig  GitlabConfig  `json:"gitlab"`
}

// JenkinsConfig is the jenkins specific config
type JenkinsConfig struct {
	URL string `json:"url"`
}

// GitlabConfig is the gitlab specific config
type GitlabConfig struct {
	URL string `json:"url"`
}
