package response

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

//Pronunciations class is pronunciations in CED entry data
type Pronunciations struct {
	Type   string   `json:"type"`
	Text   string   `json:"text"`
	Labels []string `json:"labels"`
}

//Translations class is translations in CED entry data
type Translations struct {
	Etymology int      `json:"etymology"`
	Pos       string   `json:"pos"`
	Text      string   `json:"text"`
	Labels    []string `json:"labels"`
	Note      string   `json:"note"`
}

//SuccessEntry class is CED entry data
type SuccessEntry struct {
	ID             int              `json:"id"`
	Title          string           `json:"title"`
	Digest         string           `json:"digest"`
	Pronunciations []Pronunciations `json:"pronunciations"`
	Translations   []Translations   `json:"translations"`
	Note           string           `json:"note"`
}

//DecodeSuccessEntry returns *SuccessProject instance
func DecodeSuccessEntry(r io.Reader) (*SuccessEntry, error) {
	sucData := &SuccessEntry{}
	err := json.NewDecoder(r).Decode(&sucData)
	return sucData, errors.Wrap(err, "error in response.DecodeSuccessEntry() function")
}
