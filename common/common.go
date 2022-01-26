package common

import (
	"bytes"
	"dxcserver/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"
)

func LoadConfig() model.ServerConf {
	ConfigFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read DaoxiangCity-Server Config: %s", err)
	}

	var Config model.ServerConf
	err = yaml.Unmarshal(ConfigFile, &Config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %s", err)
	}

	return Config
}

func PostJson(url string, data []byte, header []string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	req.Header.Add("content-type", "application/json")

	for i := 0; i < len(header); i += 2 {
		req.Header.Add(header[i], header[i+1])
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, _ := ioutil.ReadAll(resp.Body)

	return result, nil
}
