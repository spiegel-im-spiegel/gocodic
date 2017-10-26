package response

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

//SuccessLookup class is CED lookup data
type SuccessLookup struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Digest string `json:"digest"`
}

//DecodeSuccessLookup returns []SuccessLookup instance
func DecodeSuccessLookup(r io.Reader) ([]SuccessLookup, error) {
	sucData := make([]SuccessLookup, 0)
	err := json.NewDecoder(r).Decode(&sucData)
	return sucData, errors.Wrap(err, "error in response.DecodeSuccessLookup() function")
}
