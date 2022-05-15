package pushover

// SendMessageRequest is the type of SendMessage request messages
type SendMessageRequest struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Device  string `json:"device"`
	Title   string `json:"title"`
	Message string `json:"message"`
	HTML    string `json:"html"`
}

// SendMessageResponse is the type of SendMessage response messages
// Example: {"status":1,"request":"e4a934f5-212f-4010-b321-d9c48a244a64"}
type SendMessageResponse struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

// SendMessageErrorResponse is the type of SendMessage error response messages
// Example: {"token":"invalid","errors":["application token is invalid"],"status":0,"request":"7d9fe901-5339-4be6-8d23-0e99650ca1e2"}
type SendMessageErrorResponse struct {
	SendMessageResponse
	Token  string   `json:"token"`
	Errors []string `json:"errors"`
}
