package client

import (
	"fmt"
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
	Kind                   string                 `json:"kind"`
	Metadata               Metadata               `json:"metadata"`
	ClusterID              string                 `json:"cluster_id"`
	TopicName              string                 `json:"topic_name"`
	IsInternal             bool                   `json:"is_internal"`
	ReplicationFactor      int                    `json:"replication_factor"`
	Partitions             Partitions             `json:"partitions"`
	Configs                Configs                `json:"configs"`
	PartitionReassignments PartitionReassignments `json:"partition_reassignments"`
}

type Metadata struct {
	Self         string `json:"self"`
	ResourceName string `json:"resource_name"`
}

type Partitions struct {
	Related string `json:"related"`
}

type Configs struct {
	Related string `json:"related"`
}

type PartitionReassignments struct {
	Related string `json:"related"`
}

var client, err = NewClient()

func GetTopicConfig(topicName string) (*GetTopicResponse, error) {

	path := fmt.Sprintf("/topics/%s", topicName)
	var topic GetTopicResponse

	if err := client.Get(path, &topic); err != nil {
		fmt.Printf("Error: %q\n", err)
		return nil, err
	}
	fmt.Println(topic)
	return &topic, nil
}

func CreateTopic(requestBody PostTopicRequest) error {
	url := fmt.Sprintf("/topics/")
	if err := client.Post(url, requestBody, nil); err != nil {
		fmt.Printf("Error: %q\n", err)
		return err
	}

	fmt.Println(fmt.Sprintf("Created topic %s with following configuration %#v", requestBody.TopicName, requestBody))
	return nil
}
