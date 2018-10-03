package image_url

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		imageproxyHost string
		baseURL        string
		bucketName     string
		omitBaseURL    bool
		hashValue      string
		opts           []Option
		url            string
	}{
		{
			"http://image.test/",
			"http://minio:9000",
			"fate",
			true,
			"abcdefgheight123456789",
			[]Option{Width(100), Height(200), Quality(100)},
			"http://image.test/100x200/fate/abcdefgheight123456789",
		},
		{
			"http://image.test/",
			"http://minio:9000",
			"fate",
			false,
			"abcdefgheight123456789",
			[]Option{Width(100), Height(200), Quality(100)},
			"http://image.test/100x200/http://minio:9000/fate/abcdefgheight123456789",
		},
		{
			"http://image.test/",
			"http://minio:9000",
			"fate",
			false,
			"abcdefgheight123456789",
			[]Option{Height(200), Quality(80)},
			"http://image.test/x200,q80/http://minio:9000/fate/abcdefgheight123456789",
		},
		{
			"http://image.test/",
			"http://minio:9000",
			"fate",
			false,
			"abcdefgheight123456789",
			[]Option{HeightPercent(0.5), Quality(80)},
			"http://image.test/x0.5,q80/http://minio:9000/fate/abcdefgheight123456789",
		},
		{
			"http://image.test/",
			"http://minio:9000",
			"fate",
			false,
			"abcdefgheight123456789",
			[]Option{WidthPercent(0.6), HeightPercent(0.5), Quality(80)},
			"http://image.test/0.6x0.5,q80/http://minio:9000/fate/abcdefgheight123456789",
		},
	}
	for _, test := range tests {
		u := NewImageproxyURL(test.imageproxyHost, test.baseURL, test.bucketName, test.omitBaseURL, nil)
		v := u.Generate(test.hashValue, test.opts...)
		if test.url != v {
			t.Errorf("url生成出错, u.Generate(%s)=%s, excepted %s", test.hashValue, v, test.url)
		}
	}

}
