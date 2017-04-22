package horizon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type etConfig struct {
	token    string
	username string
	host     string
}

// Holds the users Horizon HTTP config
var config etConfig

func Get() string {
	config := getConfig()

	url := fmt.Sprintf("%s/lessons.json", config.host)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(config.username, config.token)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	bodyText, err := ioutil.ReadAll(response.Body)

	return string(bodyText)
}

func getConfig() *etConfig {
	if config != (etConfig{}) {
		return &config
	}

	token, ok := viper.Get("token").(string)

	if !ok {
		log.Fatal("No token present")
	}

	username, ok := viper.Get("username").(string)

	if !ok {
		log.Fatal("No username present")
	}

	host, ok := viper.Get("host").(string)

	if !ok {
		log.Fatal("No host present")
	}

	config := &etConfig{
		token:    token,
		username: username,
		host:     host,
	}
	return config
}
