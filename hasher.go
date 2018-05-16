package image_uploader

import (
	"io"
	"crypto/md5"
	"encoding/hex"
)

type HashFunc func(file File) (string, error)

func (hf HashFunc) Hash(file File) (string, error) {
	return hf(file)
}

type Hasher interface {
	Hash(file File) (string, error)
}

func MD5HashFunc(file File) (string, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
