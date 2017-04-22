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

func Get(path string) ([]byte, error) {
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

	body, readErr := ioutil.ReadAll(response.Body)
	return body, readErr
}

func GetLesson(slug string) *Lesson {
	b, err := Get(fmt.Sprintf("/lessons/%s.json", slug))
	if err != nil {
		log.Fatal("Failed to create url", err)
	}

	type LessonResponse struct {
		Lesson Lesson `json:"lesson"`
	}
	lr := &LessonResponse{}

	err = json.Unmarshal(b, &lr)
	if err != nil {
		log.Fatal("Failed to marshal", err)
	}
	return &lr.Lesson
}

func GetLessons() *Lessons {
	b, err := Get("lessons.json")
	if err != nil {
		log.Fatal(err)
	}

	lessons := &Lessons{}

	err = json.Unmarshal(b, &lessons)
	if err != nil {
		log.Fatal(err)
	}
	return lessons
}

func PostLesson(filepath) *Lessons {
	config := getConfig()
	url := UrlEncoded(fmt.Sprintf("%s/%s", config.host, path))

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, nil)
	request.SetBasicAuth(config.username, config.token)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() *etConfig {
	if config != (etConfig{}) {
		return &config
	}

	token := viper.GetString("token")
	username := viper.GetString("username")
	host := viper.GetString("host")

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
