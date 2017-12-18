package models

type Client struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Destinations []Destination `json:"destinations"`
	Sources      []Source      `json:"sources"`
}
