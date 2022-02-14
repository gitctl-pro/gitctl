package apps

import (
	"github.com/gitctl-pro/gitctl/pkg/model"
	"time"
)

type Cluster struct {
	ID        int       `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	Version   string    `json:"version"`
	State     string    `json:"state"`
	Type      string    `json:"type"`
	NodeCount int       `json:"nodeCount" grom:"node_count"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (a *Cluster) TableName() string {
	return "cluster"
}

type ClusterPageOption struct {
	Id        int
	Name      string
	TimeRange *model.TimeRange
	Limit     int
	Page      int
}

type Sort struct {
	Sort  string
	Order string
}

func ListCluster(query *ClusterPageOption) (apps []*Cluster, total int64) {
	tx := model.Db.Table("cluster")

	if query.TimeRange != nil {
		timeRange := query.TimeRange
		if timeRange.StartTime != "" && timeRange.EndTime != "" {
			tx = tx.Where("update_at  >= ?", timeRange.StartTime).
				Where("update_at  <= ?", timeRange.EndTime)
		}
	}
	if len(query.Name) > 0 {
		tx = tx.Where("name = ? ", query.Name)
	}

	if query.Limit == 0 {
		query.Limit = 100
	}

	tx.Count(&total)
	offset := (query.Page - 1) * query.Limit
	tx.Offset(offset).Limit(query.Limit).Order("id Desc").Scan(&apps)
	return apps, total
}

func GetCluster(query *ClusterPageOption) (c *ClusterPageOption, err error) {
	tx := model.Db.Table("application")
	if len(query.Name) > 0 {
		tx = tx.Where("name = ? ", query.Name)
	}
	if query.Id > 0 {
		tx = tx.Where("id = ? ", query.Id)
	}
	result := tx.Limit(1).Find(&c)

	return c, result.Error
}

func AddCluster(c *Cluster) error {
	c.CreateAt = time.Now()
	c.UpdateAt = time.Now()
	result := model.Db.Create(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateCluster(app *Cluster) error {
	app.UpdateAt = time.Now()
	result := model.Db.
		Where("name =?", app.Name).
		Updates(&app)
	return result.Error
}
