package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/toolkits/file"
	"log"
)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type GlobalConfig struct {
	Debug bool        `json:"debug"`
	Http  *HttpConfig `json:"http"`
}

func main() {
	var config *GlobalConfig
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()
	configContent, err := file.ToTrimString(*cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}
	config = &c
	fmt.Println(config.Http)
	d, _ := json.Marshal(config)
	fmt.Printf("%s\n", d)
}

```
cfg.json
{
    "debug": true,
    "http": {
        "enabled": true,
        "listen": "0.0.0.0:6060"
    }
}
```
