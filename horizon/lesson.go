package horizon

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
