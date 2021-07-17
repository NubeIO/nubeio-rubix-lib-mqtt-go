package mqtt_lib

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-mqtt-go/mqtt_config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type MqttConnection struct {
	mqttClient mqtt.Client
}


func NewConnection() (conn *MqttConnection) {
	//c := mqtt_config.MqttConfig("na", false)
	var br mqtt_config.Broker
	br.Host = "0.0.0.0"
	br.Port = "1883"

	c := mqtt_config.GetMqttConfig()
	opts := mqtt.NewClientOptions()
	host := "tcp://" + c.Host + ":" + c.Port
	opts.AddBroker(fmt.Sprintf(host))
	opts.SetClientID(c.ClientId)
	opts.AutoReconnect = true
	opts.OnConnectionLost = connectLostHandler
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("Connect problem: ", token.Error())
	}
	conn = &MqttConnection{client}
	return conn
}


func (con *MqttConnection) IsConnected() bool {
	connected := con.mqttClient.IsConnected()
	if !connected {
		log.Println("Healthcheck MQTT fails")
	}
	return connected
}

func (conn *MqttConnection) Publish(message string, topic string) {
	token := conn.mqttClient.Publish(topic, 1, false, message)
	token.Wait()
	log.Println("Publish to topic: ", topic)
}


var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Mqtt connected")
}

