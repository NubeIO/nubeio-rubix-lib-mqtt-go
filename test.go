package main

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-app-mqtt-go/config"
	"github.com/NubeIO/nubeio-rubix-app-mqtt-go/pkg/mqttcommon"
	"log"
)

type TMicroEdge struct {
	Sensor        string
}


func main() {


	f := "config-test.json"
	c := config.MqttConfig(f, true)
	log.Println(c)
	topic := "test"
	mqttConnection := mqttcommon.NewConnection()
	message := TMicroEdge{Sensor: "me"}
	jsonValue, _ := json.Marshal(message)
	log.Println("MQTT messages, topic:", topic, " ", "message:", message)
	mqttConnection.Publish(string(jsonValue), topic)



}
