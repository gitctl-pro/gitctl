package apps

import "time"

type Node struct {
	ID         string    `gorm:"primarykey" json:"id"`
	Cluster    string    `json:"cluster"`
	Hostname   string    `json:"hostname"`
	InternalIP string    `json:"internal_ip"`
	Status     string    `json:"status"`
	Empty      string    `json:"empty"`
	CreateAt   time.Time `json:"create_at"`
}

func (n *Node) TableName() string {
	return "k8s_node"
}
