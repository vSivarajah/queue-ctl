package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostTopicRequest struct {
	TopicName         string        `json:"topic_name"`
	PartitionsCount   int           `json:"partitions_count"`
	ReplicationFactor int           `json:"replication_factor"`
	ConfigEntries     []ConfigEntry `json:"configs,omitempty"`
}

type ConfigEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetTopicResponse struct {
	Name       string            `json:"name"`
	Config     map[string]string `json:"configs"`
	Partitions []Partition       `json:"partitions"`
}

type Partition struct {
	Partition int        `json:"partition"`
	Leader    int        `json:"leader"`
	Replicas  []Replicas `json:"replicas"`
}

type Replicas struct {
	Broker int  `json:"broker"`
	Leader bool `json:"leader"`
	InSync bool `json:"in_sync"`
}

func (c *Client) GetTopicConfig(topicName string) (*GetTopicResponse, error) {
	url := fmt.Sprintf(c.Host+"%s/topics/%s", c.ClusterId, topicName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var topic GetTopicResponse

	err = decoder.Decode(&topic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("URL:", url)

	return &topic, nil
}

func (c *Client) CreateTopic() {
	url := fmt.Sprintf(baseURL+"%s/topics/", c.ClusterId)

	requestBody := PostTopicRequest{
		TopicName:         "test-topic",
		PartitionsCount:   4,
		ReplicationFactor: 1,
		ConfigEntries: []ConfigEntry{
			ConfigEntry{
				Name:  "compression.type",
				Value: "producer",
			},
		},
	}
	json, err := json.Marshal(requestBody)
	if err != nil {
		log.Println(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Println(err)
		fmt.Println(resp.Status)
	}
}
