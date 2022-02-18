package controller

type Response struct {
	Err  error       `json:"err,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
