package config

import "gopkg.in/yaml.v2"

func LoadTopicBytes(contents []byte) (TopicConfiguration, error) {
	config := TopicConfiguration{}
	err := yaml.Unmarshal(contents, &config)
	return config, err
}
