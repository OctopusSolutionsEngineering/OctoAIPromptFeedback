package model

import "time"

type Feedback struct {
	ID        string    `jsonapi:"primary,feedback" json:"id"`
	Timestamp time.Time `jsonapi:"attribute" json:"timestamp"`
	Server    string    `jsonapi:"attr,server" json:"server"`
	Prompt    string    `jsonapi:"attribute" json:"prompt"`
	Comment   string    `jsonapi:"attribute" json:"comment"`
	ThumbsUp  bool      `jsonapi:"attribute" json:"thumbsUp"`
}
