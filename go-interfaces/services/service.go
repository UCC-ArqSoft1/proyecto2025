package services

import "go-interfaces/domain"

type Datasource interface {
	GetActivityByID(id int64) (domain.Activity, error)
}

func GetActivityByID(id int64, datasource Datasource) (domain.Activity, error) {
	return datasource.GetActivityByID(id)
}
