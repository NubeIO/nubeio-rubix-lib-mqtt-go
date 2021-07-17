package mqtt_config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var (
	B Broker
)


type Broker struct {
		Host     string
		Port     string
		Topic    string
		User     string
		Password string
		ClientId string
}


type Params struct 	{
	UseConfigFile bool
	ConfigFile    string
	GenerateFile  bool
}

//SetMqttConfig
// if args Params.GenerateFile is true this will create a json mqtt_config file and will disregard the mqtt_config Broker
// if args Params.UseConfigFile is true is will use a local mqtt_config file
func SetMqttConfig (config Broker, args Params) error {
	if !args.UseConfigFile {
		B = config
	} else {
		configFile := args.ConfigFile
		if configFile == "" {
			return errors.New("no valid mqtt_config file passed in")
		}
		file, err := os.Open(configFile)
		_config := Broker{}
		if err != nil {
			//mqtt_lib defaults
			_config.Host = "0.0.0.0"
			_config.Port = "1883"
			_config.Topic = "pub/lora"
			_config.Port = "1883"
			_config.ClientId = ""
			// generate mqtt_config file if dont exist
			if args.GenerateFile {
				j, _ := json.Marshal(_config)
				err = ioutil.WriteFile(configFile, j, 0644)
			}
			B = _config
		} else {
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&_config)
			B = _config
		}
	}
	return nil

}

func GetMqttConfig()  Broker {
	return B
}
