package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"hfs_backend/common"
	"hfs_backend/component/uploadprovider"
	"hfs_backend/modules/upload/uploadmodel"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, image *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, storage CreateImageStorage) *uploadBiz {
	return &uploadBiz{
		provider: provider,
		imgStore: storage,
	}
}
func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, filename string) (*common.Image, error) {
	fileByte := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileByte)
	if err != nil {
		return nil, uploadmodel.ErrFileIsnotImage(err)
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	fileExt := filepath.Ext(filename)
	fileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)
	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil

}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Print("err: ", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
