package main

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/mqtt_config"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"log"
)

type Mes struct {
	M        string `json:"my_cool_key"`
}


func main() {
	var mqtt mqtt_config.Params
	mqtt.UseConfigFile = false

	var br mqtt_config.Broker
	br.Host = "0.0.0.0"
	br.Port = "1883"
	err := mqtt_config.SetMqttConfig(br, mqtt); if err != nil {
		log.Println(err)
		return
	}
	topic := "test"
	mqttConnection := mqtt_lib.NewConnection()
	message := Mes{M: "me"}
	jsonValue, _ := json.Marshal(message)
	log.Println("MQTT messages, topic:", topic, " ", "message:", message)
	mqttConnection.Publish(string(jsonValue), topic)



}
