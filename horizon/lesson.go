package horizon

type Lesson struct {
	Slug  string `json:"slug,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Lessons struct {
	Lessons []Lesson `json:"lessons"`
}
