package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vsivarajah/queue-ctl/pkg/config"
)

type conf struct {
	Hits     int64 `yaml:"hits" json:"hits,omitempty"`
	Time     int64 `yaml:"time" json:"time,omitempty"`
	Deadline int64 `yaml:"deadline,omitempty"`
}

func main() {

	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	topicConfig, err := config.LoadTopicBytes(yamlFile)
	if err != nil {
		log.Println(err)
	}
	mapConfig := []config.ConfigEntry{}
	for k, v := range topicConfig.Spec.ConfigEntries {
		mapConfig = append(mapConfig, config.ConfigEntry{
			Name:  k,
			Value: v,
		})
	}

	fmt.Println(mapConfig[0].Name)
}
