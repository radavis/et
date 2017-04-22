package horizon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type etConfig struct {
	token    string
	username string
	host     string
}

// Holds the users Horizon HTTP config
var config etConfig

func Get(path string) string {
	config := getConfig()
	url := UrlEncoded(fmt.Sprintf("%s/%s", config.host, path))

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(config.username, config.token)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	bodyText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyText)
}

func GetLessons() *Lessons {
	lessonJson := Get("lessons.json")
	lessons := &Lessons{}

	err := json.Unmarshal([]byte(lessonJson), &lessons)
	if err != nil {
		log.Fatal(err)
	}

	return lessons
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

// UrlEncoded encodes a string like Javascript's encodeURIComponent()
func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		log.Fatal(err)
	}
	return u.String()
}
