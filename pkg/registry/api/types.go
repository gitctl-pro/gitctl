package api

import "time"

const (
	ACTION_PULL string = "pull"
	ACTION_PUSH        = "push"
)

type RecordEvents struct {
	Events []*Event `json:"events"`
}

type actor struct {
	Name string `json:"name"`
}

type request struct {
	Addr      string `json:"addr"`
	Host      string `json:"host"`
	ID        string `json:"id"`
	Method    string `json:"method"`
	UserAgent string `json:"useragent"`
}

type source struct {
	Addr       string `json:"addr"`
	InstanceID string `json:"instanceID"`
}

type target struct {
	Digest     string `json:"digest"`
	Length     int    `json:"length"`
	MediaType  string `json:"mediaType"`
	Repository string `json:"repository"`
	Size       int    `json:"size"`
	Tag        string `json:"tag"`
	URL        string `json:"url"`
}

type Event struct {
	ID        string    `json:"id"`
	Action    string    `json:"action"`
	Actor     *actor    `json:"actor"`
	Request   *request  `json:"request"`
	Source    *source   `json:"source"`
	Target    *target   `json:"target"`
	Timestamp time.Time `json:"timestamp"`
}
