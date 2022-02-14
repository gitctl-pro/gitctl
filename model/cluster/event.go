package apps

import "time"

type Event struct {
	ID              string    `gorm:"primarykey" json:"id"`
	Cluster         string    `json:"cluster"`
	Namespace       string    `json:"namespace"`
	Kind            string    `json:"kind" `
	Name            string    `json:"name"`
	Reason          string    `json:"reason"`
	Type            string    `json:"type"`
	Message         string    `json:"message"`
	SourceComponent string    `json:"source_component"`
	SourceHost      string    `json:"source_host"`
	FirstTime       time.Time `json:"first_time"`
	LastTime        time.Time `json:"last_time"`
}

func (e *Event) TableName() string {
	return "k8s_event"
}
