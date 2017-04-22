package horizon

import (
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "github.com/spf13/viper"
)

func Get() string {
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

    url := fmt.Sprintf("%s/lessons.json", host)

    client := &http.Client{}
    request, err := http.NewRequest("GET", url, nil)
    request.SetBasicAuth(username, token)

    response, err := client.Do(request)
    if err != nil{
        log.Fatal(err)
    }

    defer response.Body.Close()
    bodyText, err := ioutil.ReadAll(response.Body)

    return string(bodyText)
}