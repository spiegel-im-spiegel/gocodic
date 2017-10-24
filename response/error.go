package response

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

//ErrorData class is body data at error
type ErrorData struct {
	Message string `json:"message"`
	Context string `json:"context"`
	Code    int    `json:"code"`
}

//ErrorList class is list of ErrorData
type ErrorList struct {
	Errors []ErrorData `json:"errors"`
}

//DecodeError returns ErrorData instance
func DecodeError(r io.Reader) (*ErrorList, error) {
	errData := &ErrorList{}
	err := json.NewDecoder(r).Decode(errData)
	return errData, errors.Wrap(err, "error in response.DecodeError() function")
}
