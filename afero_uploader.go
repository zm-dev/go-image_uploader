package image_uploader

import (
	"github.com/spf13/afero"
	"io"
	"path/filepath"
)

type aferoUploader struct {
	h  Hasher
	s  Store
	fs afero.Fs
}

func (au *aferoUploader) saveToFs(hashValue string, f File) error {
	_, err := f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	// todo savepath
	newFile, err := au.fs.Create(hashValue)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, f)
	return err
}

func (au *aferoUploader) Upload(fh FileHeader) (*Image, error) {
	info, err := DecodeImageInfo(fh.File)
	if err != nil {
		return nil, err
	}

	hashValue, err := au.h.Hash(fh.File)
	if err != nil {
		return nil, err
	}

	if exist, err := au.s.ImageExist(hashValue); exist && err == nil {
		// 图片已经存在
		return au.s.ImageLoad(hashValue)
	} else if err != nil {
		return nil, err
	}

	if err := au.saveToFs(hashValue, fh.File); err != nil {
		return nil, err
	}
	return saveToStore(au.s, hashValue, fh.Filename, info)
}

func (au *aferoUploader) UploadFromURL(u string, filename string) (*Image, error) {
	if filename != "" {
		filename = filepath.Base(u)
	}
	file, size, err := DownloadImage(u)

	if err != nil {
		return nil, err
	}

	defer removeFile(file)

	fh := FileHeader{
		Filename: filename,
		Size:     size,
		File:     file,
	}
	return au.Upload(fh)
}

func NewAferoUploader(h Hasher, s Store, fs afero.Fs) Uploader {
	return &aferoUploader{h: h, s: s, fs: fs}
}
