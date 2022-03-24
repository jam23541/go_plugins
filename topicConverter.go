package go_plugins

import "strings"

func GetMqttTopic(natsTopic string) string {
	return strings.ReplaceAll(natsTopic, ".", "/")
}

func GetNatsTopic(mqttTopic string) string {
	return strings.ReplaceAll(mqttTopic, "/", ".")
}
