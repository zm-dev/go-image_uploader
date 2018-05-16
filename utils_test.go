package image_uploader

import (
	"testing"
	"os"
)

func TestDecodeImageInfo(t *testing.T) {
	tests := []struct {
		filename string
		info     ImageInfo
	}{
		{"testdata/Go-Logo_Aqua.jpg", ImageInfo{1062, 938, "jpeg"}},
		{"testdata/Go-Logo_Black.jpg", ImageInfo{1061, 938, "jpeg"}},
		{"testdata/Go-Logo_Blue.jpg", ImageInfo{1061, 938, "jpeg"}},
		{"testdata/gopher.png", ImageInfo{200, 200, "png"}},
	}

	for _, test := range tests {
		file, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}

		info, err := DecodeImageInfo(file)
		if err != nil {
			t.Errorf("unexpected error. %+v", err)
		}
		if info != test.info {
			t.Errorf("DecodeImageInfo(%s)=%+v ,expected %+v", test.filename, info, test.info)
		}

		file.Close()
	}
}

func TestDecodeImageInfoOnUnsupportFormat(t *testing.T) {
	tests := []struct {
		filename string
	}{
		{"testdata/test.txt"},
	}

	for _, test := range tests {
		file, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}

		_, err = DecodeImageInfo(file)

		if !IsUnknownFormat(err) {
			t.Error(err)
			t.Errorf("excepted ErrFormat")
		}

		file.Close()
	}
}
