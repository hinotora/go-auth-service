package models

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}
