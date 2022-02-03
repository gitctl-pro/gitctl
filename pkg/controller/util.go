package controller

type Response struct {
	Err  error       `json:"err,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
