package image_uploader

import (
	"testing"
	"os"
)

func TestMD5HashFunc(t *testing.T) {
	tests := []struct {
		filename  string
		hashValue string
	}{
		{"testdata/Go-Logo_Aqua.jpg", "6e511bc79b63156646f2094946d10c59"},
		{"testdata/Go-Logo_Black.jpg", "6ed618ff84042c827ba8dc7552f24fc1"},
		{"testdata/Go-Logo_Blue.jpg", "b76dac1fbbe373cae8927f8a91ece665"},
		{"testdata/Go-Logo_Fuchsia.jpg", "5d9a512547df2340adad8ed588a417d0"},
		{"testdata/Go-Logo_LightBlue.jpg", "fd8b1d1249ecbad1595c60270907b5de"},
		{"testdata/Go-Logo_Yellow.jpg", "dbd7f8f9f43efb0f830673ca7f6130d7"},
		{"testdata/gopher.png", "1faaf9020e0df18fdf0429e0db211f37"},
		{"testdata/Go-Logo_Aqua.jpg", "6e511bc79b63156646f2094946d10c59"},
		{"testdata/Go-Logo_Black.jpg", "6ed618ff84042c827ba8dc7552f24fc1"},
		{"testdata/Go-Logo_Blue.jpg", "b76dac1fbbe373cae8927f8a91ece665"},
		{"testdata/Go-Logo_Fuchsia.jpg", "5d9a512547df2340adad8ed588a417d0"},
		{"testdata/Go-Logo_LightBlue.jpg", "fd8b1d1249ecbad1595c60270907b5de"},
		{"testdata/Go-Logo_Yellow.jpg", "dbd7f8f9f43efb0f830673ca7f6130d7"},
		{"testdata/gopher.png", "1faaf9020e0df18fdf0429e0db211f37"},
	}

	for _, test := range tests {
		file, err := os.Open(test.filename)
		if err != nil {
			t.Error(err)
		}
		hashValue, err := MD5HashFunc(file)
		if err != nil {
			t.Errorf("unexpected error. error: %+v", err)
		}

		if hashValue != test.hashValue {
			t.Errorf("MD5HashFunc(%s) = %s, excepted %s", test.filename, hashValue, test.hashValue)
		}
		file.Close()
	}
}
