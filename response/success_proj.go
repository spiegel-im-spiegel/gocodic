package response

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

//Owner class is owner data
type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//SuccessProject class is project data
type SuccessProject struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	WordsCountOn int    `json:"words_count"`
	CreateOn     string `json:"create_on"`
	Owner        Owner  `json:"owner"`
}

//DecodeSuccessProject returns *SuccessProject instance
func DecodeSuccessProject(r io.Reader) (*SuccessProject, error) {
	sucData := &SuccessProject{}
	err := json.NewDecoder(r).Decode(&sucData)
	return sucData, errors.Wrap(err, "error in response.DecodeSuccessProject() function")
}

//DecodeSuccessProjects returns []SuccessProject instance
func DecodeSuccessProjects(r io.Reader) ([]SuccessProject, error) {
	sucData := make([]SuccessProject, 0)
	err := json.NewDecoder(r).Decode(&sucData)
	return sucData, errors.Wrap(err, "error in response.DecodeSuccessProjects() function")
}
