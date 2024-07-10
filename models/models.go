package models

import "time"

type Config struct {
	Website      string
	Author       string
	Description  string
	Title        string
	PostsOnIndex int `toml:"postsOnIndex"`
}

type Post struct {
	Title       string
	Description string
	Date        time.Time
	Draft       bool
	Content     string
	Slug        string
	CSS         string `toml:"css"`
}
