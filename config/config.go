package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)


type Configuration struct {
	Broker struct {
		Host     string
		Port     string
		Topic    string
		User     string
		Password string
		ClientId string
	}
}

func MqttConfig(configFile string, generateFile bool) Configuration {
	file, err := os.Open(configFile)
	Config := Configuration{}
	if err != nil {
		//mqtt defaults
		Config.Broker.Host = "0.0.0.0"
		Config.Broker.Port = "1883"
		Config.Broker.Topic = "pub/lora"
		Config.Broker.Port = "1883"
		Config.Broker.ClientId = ""
		// generate config file if dont exist
		if generateFile {
			j, _ := json.Marshal(Config)
			err = ioutil.WriteFile(configFile, j, 0644)
		}
		return Config
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
		return Config
	}

}

