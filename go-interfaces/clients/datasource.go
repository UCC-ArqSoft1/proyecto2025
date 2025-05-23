package clients

import (
	"go-interfaces/domain"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

// service
func (mySQLDatasource MySQL) GetActivityByID(id int64) (domain.Activity, error) {
	var activity domain.Activity

	result := mySQLDatasource.DB.First(&activity, id)
	if result.Error != nil {
		return domain.Activity{}, result.Error
	}
	return activity, nil
}

type SQLite struct {
	Connection string
}

func (SQLite) GetActivityByID(id int64) (domain.Activity, error) {
	return domain.Activity{}, nil
}
