package image_uploader

import (
	"testing"
	"github.com/dmgk/faker"
	"time"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"log"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func openTest() *gorm.DB {
	return open("sqlite3", "file::memory:?cache=shared")
}

func open(driverName, dataSourceName string) (*gorm.DB) {
	db, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
		log.Fatalln("database connection failed")
	}
	return db
}

// var testTx *gorm.DB
//func TestMain(m *testing.M) {
//
//	testTx = openTest()
//	image := &model.Image{}
//	if !testTx.HasTable(image) {
//		testTx.CreateTable(image)
//	}
//	testTx.Begin()
//	defer testTx.Rollback()
//	m.Run()
//}

var formats = []string{
	"jpeg", "png", "gif",
}

func mockImage() *Image {
	var (
		createdAt = faker.Time().Backward(1000 * time.Hour)
		updatedAt = createdAt.Add(time.Duration(faker.RandomInt(0, 1000)) * time.Hour)
	)
	return &Image{
		Hash: faker.RandomString(32),
		// Url:       faker.Avatar().String(),
		Format:    faker.RandomChoice(formats),
		Title:     faker.Name().Title(),
		Width:     uint(faker.RandomInt(1, 2000)),
		Height:    uint(faker.RandomInt(1, 2000)),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func setUpDBForTesting() *gorm.DB {
	testTx := openTest()
	image := &Image{}
	if !testTx.HasTable(image) {
		testTx.CreateTable(image)
	}
	return testTx.Begin()
}

func TestDbImageStore_Create(t *testing.T) {
	testTx := setUpDBForTesting()
	defer func() {
		testTx.Rollback().Close()
	}()

	is := NewDBStore(testTx)
	image := mockImage()
	err := is.ImageCreate(image)
	if err != nil {
		t.Errorf("unexpected error. error: %+v", err)
	}

	err = is.ImageCreate(image)

	if err == nil {
		t.Error("hash 主键不能重复")
	}
}

func TestDbImageStore_Load(t *testing.T) {
	testTx := setUpDBForTesting()
	defer testTx.Rollback()

	image := mockImage()
	is := NewDBStore(testTx)
	// 这个时候还没有数据
	_, err := is.ImageLoad(image.Hash)
	if err == nil {
		t.Error("hash应该不存在")
	}
	is.ImageCreate(image)
	loadImage, err := is.ImageLoad(image.Hash)
	if err != nil {
		t.Errorf("unexpected error. error: %+v", err)
	}
	b1, _ := json.Marshal(image)
	b2, _ := json.Marshal(loadImage)
	if string(b1) != string(b2) {
		t.Error("loadImage != image")
	}
}
