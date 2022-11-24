package domains

type Message struct {
	Target  string `json:"target"`
	Content string `json:"content"`
}
