package horizon

import (
	"fmt"
	"os"
	"net/http"
	"io"
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
	output_file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error while creating", filename, "-", err)
		return
	}
	defer output_file.Close()

	response, err := http.Get(l.ArchiveUrl)
	if err != nil {
		fmt.Println("Error while downloading", l.ArchiveUrl, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output_file, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", l.ArchiveUrl, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}
