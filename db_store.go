package image_uploader

import (
	"github.com/jinzhu/gorm"
)

type dbStore struct {
	db *gorm.DB
}

func (is *dbStore) ImageLoad(hash string) (image *Image, err error) {
	image = &Image{}
	err = is.db.Where(Image{Hash: hash}).First(image).Error
	return
}

func (is *dbStore) ImageCreate(image *Image) error {
	return is.db.Create(image).Error
}

func (is *dbStore) ImageExist(hash string) (bool, error) {
	var count uint
	err := is.db.Model(&Image{}).Where(Image{Hash: hash}).Count(&count).Error
	return count > 0, err
}

func NewDBStore(db *gorm.DB) Store {
	return &dbStore{db}
}
