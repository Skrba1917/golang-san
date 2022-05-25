package main

type Config struct {
	Entries map[string]string `json:"entries"`
	Id      string            `json:"id"`
	Label   string            `json:"label"`
	Version string            `json:"version"`
}

type Group struct {
	Entries []map[string]string `json:"entries"`
	Id      string              `json:"id"`
	Version string              `json:"version"`
}
