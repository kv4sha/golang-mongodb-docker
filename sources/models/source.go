package models

type Source struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	SourceType string      `json:"sourceType"`
	Setting    interface{} `json:"setting"`
}
