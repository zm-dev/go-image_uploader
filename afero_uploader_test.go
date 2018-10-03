package image_uploader
//
//import (
//	"testing"
//	"github.com/spf13/afero"
//	"os"
//	"path"
//)
//
//func TestAferoUploader_Upload(t *testing.T) {
//	testTx := setUpDBForTesting()
//	defer testTx.Rollback()
//	fs := afero.NewMemMapFs()
//	store := NewDBStore(testTx)
//	uploader := NewAferoUploader(HashFunc(MD5HashFunc), store, fs, nil)
//	tests := []struct {
//		filename string
//	}{
//		{"testdata/Go-Logo_Aqua.jpg"},
//		{"testdata/Go-Logo_Black.jpg"},
//		{"testdata/Go-Logo_Blue.jpg"},
//		{"testdata/Go-Logo_Fuchsia.jpg"},
//		{"testdata/Go-Logo_LightBlue.jpg"},
//		{"testdata/Go-Logo_Yellow.jpg"},
//		{"testdata/gopher.png"},
//	}
//
//	for _, test := range tests {
//		basename := path.Base(test.filename)
//		file, err := os.Open(test.filename)
//		if err != nil {
//			t.Error(err)
//		}
//		image, err := uploader.Upload(FileHeader{File: file, Filename: basename})
//		if err != nil {
//			t.Error(err)
//		}
//		file.Close()
//		if _, err := fs.Stat(image.Hash); err != nil {
//			if os.IsNotExist(err) {
//				t.Error("文件保存失败")
//			} else {
//				t.Error(err)
//			}
//		}
//
//		if exist, err := store.ImageExist(image.Hash); err != nil {
//			t.Error(err)
//		} else if !exist {
//			t.Error("图片保存到store中失败")
//		}
//	}
//
//}
//
//func TestAferoUploader_UploadFromURL(t *testing.T) {
//	testTx := setUpDBForTesting()
//	defer testTx.Rollback()
//	fs := afero.NewMemMapFs()
//	store := NewDBStore(testTx)
//	uploader := NewAferoUploader(HashFunc(MD5HashFunc), store, fs, nil)
//	tests := []struct {
//		url string
//	}{
//		{"https://gss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2017-09-27/297f5edb1e984613083a2d3cc0c5bb36.png"},
//	}
//
//	for _, test := range tests {
//		image, err := uploader.UploadFromURL(test.url, "")
//		if err != nil {
//			t.Error(err)
//		}
//		if _, err := fs.Stat(image.Hash); err != nil {
//			if os.IsNotExist(err) {
//				t.Error("文件保存失败")
//			} else {
//				t.Error(err)
//			}
//		}
//		if exist, err := store.ImageExist(image.Hash); err != nil {
//			t.Error(err)
//		} else if !exist {
//			t.Error("图片保存到store中失败")
//		}
//	}
//}
