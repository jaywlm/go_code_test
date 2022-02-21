package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
	Mysql
	Cache
}
type Mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type Cache struct {
	Enable bool     `yaml:"enable"`
	List   []string `yaml:"list,flow"`
}

func main() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Println(err)
	}
	resultMap := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
	fmt.Println(resultMap)
}

```
# cat test.yaml 
cache:
  enable : false
  list : [redis,mongoDB]
mysql:
  user : a111
  password : 1111
  host : 1.1.2.3
  port : 3306
  name : aaa

```
