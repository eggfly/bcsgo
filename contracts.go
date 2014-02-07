package bcsgo

type ObjectCollection struct {
	ObjectTotal int       `json:"object_total"`
	Start       int       `json:"start"`
	Limit       int       `json:"limit"`
	Bucket      string    `json:"bucket"`
	Objects     []*Object `json:"object_list"`
}
