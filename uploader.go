package image_uploader

import (
	"io"
	"context"
	"errors"
)

type File interface {
	io.Reader
	io.Seeker
}

type FileHeader struct {
	Filename string
	Size     int64
	File     File
}

type Uploader interface {
	Upload(fh FileHeader) (*Image, error)
}

func Upload(ctx context.Context, fh FileHeader) (*Image, error) {
	u, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("uploader不存在")
	}
	return u.Upload(fh)
}
