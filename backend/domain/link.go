package domain

import (
	"crud-api-class/infra"
)

type Url struct {
	ID   uint64 `json:"id"`
	Link string `json:"url"`
}

func GetAll() ([]*Url, error) {
	var l []*Url

	db := infra.Connect()
	defer db.Close()

	err := db.Find(&l).Error
	if err != nil {
		return nil, err
	}

	return l, err
}

func CreateNewElement(url string) error {
	newURL := Url{}
	newURL.Link = url

	db := infra.Connect()
	defer db.Close()

	return db.Create(&newURL).Error
}

func UpdateElement(id uint64, url string) error {
	db := infra.Connect()
	defer db.Close()

	newURL := Url{}

	return db.Where("id = ?", id).First(&newURL).Update("link", url).Error
}
