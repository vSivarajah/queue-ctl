package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vsivarajah/queue-ctl/pkg/client"
	"github.com/vsivarajah/queue-ctl/pkg/config"
)

func main() {

	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	topicConfig, err := config.LoadTopicBytes(yamlFile)
	if err != nil {
		log.Println(err)
	}
	mapConfig := []client.ConfigEntry{}
	for k, v := range topicConfig.Spec.ConfigEntries {
		mapConfig = append(mapConfig, client.ConfigEntry{
			Name:  k,
			Value: v,
		})
	}

	requestBody := client.PostTopicRequest{
		TopicName:         topicConfig.Spec.TopicName,
		PartitionsCount:   topicConfig.Spec.PartitionsCount,
		ReplicationFactor: topicConfig.Spec.ReplicationFactor,
		ConfigEntries:     mapConfig,
	}
	client.CreateTopic(requestBody)

	topic, err := client.GetTopicConfig("luuulia")
	if err != nil {

		log.Println("Error:", err)

	}
	fmt.Println(topic.TopicName)

}
