package image_url

import (
	"context"
	"errors"
)

var (
	XSOptions = []Option{Width(50), Height(50), Quality(90)}
	SMOptions = []Option{Width(100), Height(100), Quality(90)}
	MDOptions = []Option{Width(200), Height(200), Quality(90)}
	LGOptions = []Option{Width(400), Height(400), Quality(90)}
)

type options struct {
	width, height               uint
	widthPercent, heightPercent float32
	quality                     uint8
}

type Option func(*options)

func Width(width uint) Option {
	return func(o *options) {
		o.width = width
	}
}

func Height(height uint) Option {
	return func(o *options) {
		o.height = height
	}
}

func WidthPercent(widthPercent float32) Option {
	return func(o *options) {
		o.widthPercent = widthPercent
	}
}

func HeightPercent(heightPercent float32) Option {
	return func(o *options) {
		o.heightPercent = heightPercent
	}
}

func Quality(quality uint8) Option {
	return func(o *options) {
		o.quality = quality
	}
}

var defaultURLOptions = options{
	quality: 90,
}

type URL interface {
	Generate(hashValue string, opt ...Option) string
}

func Generate(ctx context.Context, hashValue string, opt ...Option) (string, error) {
	url, ok := FromContext(ctx)
	if !ok {

		return "", errors.New("context中不存在 URL")
	}
	return url.Generate(hashValue, opt...), nil
}
