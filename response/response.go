package response

import (
	"bytes"
	"io"
	"net/http"
)

//Response class is response data from codic service
type Response struct {
	statusCode int
	status     string
	body       []byte
}

//New returns Response instance
func New(statusCode int, status string, body []byte) *Response {
	return &Response{statusCode: statusCode, status: status, body: body}
}

//StatusCode returns status code in response
func (r *Response) StatusCode() int {
	return r.statusCode
}

//IsSuccess returns boolean of checking status
func (r *Response) IsSuccess() bool {
	return r.statusCode != 0 && r.statusCode < http.StatusBadRequest
}

//Status returns status (string) in response
func (r *Response) Status() string {
	return r.status
}

//Body returns body data (io.Reader) in response
func (r *Response) Body() io.Reader {
	return bytes.NewReader(r.body)
}

//Stringer returns body data (string) in response
func (r *Response) String() string {
	return string(r.body)
}
