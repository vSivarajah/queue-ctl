package config

type TopicConfiguration struct {
	Spec TopicSpecification `yaml:"spec"`
}

type TopicSpecification struct {
	TopicName         string            `yaml:"topic_name"`
	PartitionsCount   int               `yaml:"partitions_count"`
	ReplicationFactor int               `yaml:"replication_factor"`
	ConfigEntries     map[string]string `yaml:"configs,omitempty"`
}

type ConfigEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
