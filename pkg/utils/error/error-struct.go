package error

// Response is generic struct with message and status
// code fields. It is used for responding with nicely formatted
// error messages
type Response struct {
	Msg        string `json:"msg,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

// Back is generic function that returns error embedded
// in struct. It requires http status code and corresponding message
func Back(sc int, msg string) *Response {
	return &Response{
		Msg:        msg,
		StatusCode: sc,
	}
}
