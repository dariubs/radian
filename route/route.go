package route

import (
	// config
	"github.com/BurntSushi/toml"
	"github.com/dariubs/radian/config"

	// builtin
	"log"
)

var err error
var Config config.CONFIG

type RESPONSE struct {
	Data  DATA  `json:"data"`
	Error ERROR `json:"error"`
}

type DATA struct {
	Filename string `json:"filename"`
	Message  string `json:"message"`
	Url      string `json:"url"`
}

type ERROR struct {
	HasError     bool   `json:"has_error"`
	ErrorNumber  int    `json:"error_number"`
	ErrorMessage string `json:"error_message"`
}

func init() {
	if _, err = toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
		return
	}
}
