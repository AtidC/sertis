package card

type Response struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
	Result    interface{} `json:"result,omitempty"`
}

type CardRequest struct {
	Id       string `json:"id"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Content  string `json:"content"`
	AuthorID string
}

type CardResponse struct {
	RequestId string `json:"requestId"`
}

type Card struct {
	Id         string `json:"id"`
	Category   string `json:"category"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	CreateDate string `json:"create_timestamp"`
	UpdateDate string `json:"update_timestamp"`
}
