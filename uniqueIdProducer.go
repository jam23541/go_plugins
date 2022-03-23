package go_plugins

import (
	"strings"
	"time"
)

func UniqueId(puber string) string {
	return puber + time.Now().Format(time.StampMilli)
}

func GetMqttTopic(natsTopic string) string {
	return strings.ReplaceAll(natsTopic, ".", "/")
}

func GetNatsTopic(mqttTopic string) string {
	return strings.ReplaceAll(mqttTopic, "/", ".")
}
