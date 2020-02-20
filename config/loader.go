package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Loader struct{}

func (loader Loader) Load(handle io.Reader) (conf Config, err error) {
	data, _ := ioutil.ReadAll(handle)
	json.Unmarshal(data, &conf)
	return
}
