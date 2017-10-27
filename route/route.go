package route

import(
  // config
	"github.com/BurntSushi/toml"
	"github.com/dariubs/radian/config"

  // builtin
  "log"
)

var err error
var Config config.CONFIG

func init(){
  if _, err = toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
		return
	}
}
