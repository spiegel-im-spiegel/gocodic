package options

import "strconv"

//ProjectID class is "project_id" parameter
type ProjectID int

//Key returns key string
func (pid ProjectID) Key() string {
	return "project_id"
}

//Value returns key string
func (pid ProjectID) Value() string {
	return strconv.Itoa(int(pid))
}
