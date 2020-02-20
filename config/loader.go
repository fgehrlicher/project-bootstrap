package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Loader struct{}

func (loader Loader) Load(handle io.Reader) (conf Config, err error) {
	data, err := ioutil.ReadAll(handle)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &conf)
	return
}
