package response

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

//Text class is text data
type Text struct {
	Text string `json:"text"`
}

//Word class is translated word information
type Word struct {
	Successful     bool   `json:"successful"`
	Text           string `json:"text"`
	TranslatedText string `json:"translated_text"`
	Candidates     []Text `json:"candidates"`
}

//SuccessTrans class is body data at error
type SuccessTrans struct {
	Successful     bool   `json:"successful"`
	Text           string `json:"text"`
	TranslatedText string `json:"translated_text"`
	Words          []Word `json:"words"`
}

//DecodeSuccessTrans returns []SuccessTrans instance
func DecodeSuccessTrans(r io.Reader) ([]SuccessTrans, error) {
	sucData := make([]SuccessTrans, 0)
	err := json.NewDecoder(r).Decode(&sucData)
	return sucData, errors.Wrap(err, "error in response.DecodeSuccess() function")
}
