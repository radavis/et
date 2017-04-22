package horizon

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"io/ioutil"

	"github.com/mholt/archiver"
)

type Lesson struct {
	Slug       string `json:"slug,omitempty"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	ArchiveUrl string `json:"archive_url,omitempty"`
	Body       string `json:"body,omitempty"`
}

type Lessons struct {
	Lessons []Lesson `json:"lessons"`
}

func (l *Lesson) GetArchive() {
	filename := fmt.Sprintf("%s.tar.gz", l.Slug)

	tempfile, err := ioutil.TempFile("", filename)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempfile.Name())

	response, err := http.Get(l.ArchiveUrl)
	if err != nil {
		fmt.Println("Error while downloading", l.ArchiveUrl, "-", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tempfile.Write(body);
	if err != nil {
		log.Fatal(err)
	}

	err = archiver.TarGz.Open(tempfile.Name(), ".")
	if err != nil {
		log.Fatal(err)
	}
}
