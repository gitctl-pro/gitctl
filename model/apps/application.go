package apps

import "time"

type Application struct {
	ID        string    `gorm:"primarykey" json:"id"`
	Cluster   string    `json:"cluster"`
	Namespace string    `json:"namespace"`
	Kind      string    `json:"kind" `
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (a *Application) TableName() string {
	return "k8s_application"
}
