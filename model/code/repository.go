package code

import (
	"github.com/gitctl-pro/gitctl/pkg/model"
	"time"
)

type Repository struct {
	ID        int       `gorm:"primarykey" json:"id"`
	Cluster   string    `json:"cluster"`
	Namespace string    `json:"namespace"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	Status    string    `json:"status"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (a *Repository) TableName() string {
	return "code_repository"
}

type RepoPageOption struct {
	Namespace string
	Name      string
	Id        int
	TimeRange *model.TimeRange
	Limit     int
	Page      int
}

func ListRepository(query *RepoPageOption) (apps []*Repository, total int64) {
	tx := model.Db.Table("code_repository")

	if query.TimeRange != nil {
		timeRange := query.TimeRange
		if timeRange.StartTime != "" && timeRange.EndTime != "" {
			tx = tx.Where("update_at  >= ?", timeRange.StartTime).
				Where("update_at  <= ?", timeRange.EndTime)
		}
	}

	if len(query.Namespace) > 0 {
		tx = tx.Where("namespace = ?", query.Namespace)
	}

	if len(query.Name) > 0 {
		tx = tx.Where("name = ?", query.Name)
	}

	if query.Limit == 0 {
		query.Limit = 100
	}

	tx.Count(&total)
	offset := (query.Page - 1) * query.Limit
	tx.Offset(offset).Limit(query.Limit).Order("id Desc").Scan(&apps)
	return apps, total
}

func GetRepository(query *RepoPageOption) (app *Repository, err error) {
	tx := model.Db.Table("code_repository")
	if len(query.Name) > 0 {
		tx = tx.Where("name = ?", query.Name)
	}
	if query.Id > 0 {
		tx = tx.Where("id = ?", query.Id)
	}
	result := tx.Limit(1).Find(&app)

	return app, result.Error
}

func AddRepository(app *Repository) error {
	app.CreateAt = time.Now()
	app.UpdateAt = time.Now()
	result := model.Db.Create(&app)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRepository(app *Repository) error {
	app.UpdateAt = time.Now()
	result := model.Db.Updates(&app).Where("name =?", app.Name).Where("namespace =?", app.Namespace)
	return result.Error
}
