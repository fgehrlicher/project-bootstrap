package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Loader can be used to retrieve configs
type Loader struct{}

// Load returns a validated conf struct for a given reader handle
func (loader Loader) Load(handle io.Reader) (conf Config, err error) {
	data, err := ioutil.ReadAll(handle)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &conf)
	return
}
