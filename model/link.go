package model

type Link struct {
	Source int64 `json:"source"`
	Target int64 `json:"target"`
	Cost   int64 `json:"cost"`
}
