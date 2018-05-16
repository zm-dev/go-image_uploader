package image_uploader

import (
	"testing"
	"github.com/spf13/afero"
	"os"
	"path"
)

func TestAferoUploader_Upload(t *testing.T) {
	testTx := setUpDBForTesting()
	defer testTx.Rollback()
	fs := afero.NewMemMapFs()
	store := NewDBStore(testTx)
	uploader := NewAferoUploader(HashFunc(MD5HashFunc), store, fs)
	tests := []struct {
		filename string
	}{
		{"testdata/Go-Logo_Aqua.jpg"},
		{"testdata/Go-Logo_Black.jpg"},
		{"testdata/Go-Logo_Blue.jpg"},
		{"testdata/Go-Logo_Fuchsia.jpg"},
		{"testdata/Go-Logo_LightBlue.jpg"},
		{"testdata/Go-Logo_Yellow.jpg"},
		{"testdata/gopher.png"},
	}

	for _, test := range tests {
		basename := path.Base(test.filename)
		file, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}
		image, err := uploader.Upload(FileHeader{File: file, Filename: basename})
		if err != nil {
			t.Error(err)
		}
		file.Close()
		if _, err := fs.Stat(image.Hash); err != nil {
			if os.IsNotExist(err) {
				t.Error("文件保存失败")
			} else {
				t.Error(err)
			}
		}

		if exist, err := store.ImageExist(image.Hash); err != nil {
			t.Error(err)
		} else if !exist {
			t.Error("图片保存到store中失败")
		}
	}

}
