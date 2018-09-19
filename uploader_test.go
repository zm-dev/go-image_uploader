package image_uploader

import (
	"testing"
	"io/ioutil"
	"io"
)

func TestDownloadImage(t *testing.T) {
	tests := []struct {
		url string
	}{
		{"https://gss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2017-09-27/297f5edb1e984613083a2d3cc0c5bb36.png"},
	}

	for _, test := range tests {
		f, size, err := DownloadImage(test.url)
		if err != nil {
			t.Error(err)
		}
		f.Seek(0, io.SeekStart)
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Error(err)
		}
		removeFile(f)
		if len(b) != int(size) {
			t.Errorf("len(%d) != size(%d)", len(b), int(size))
		}
	}
}
