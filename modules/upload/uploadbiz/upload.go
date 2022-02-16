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
	w, h, error := getImageDimension(fileByte)
	if error != nil {
		return nil, uploadmodel.ErrFileIsnotImage(error)
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	fileExt := filepath.Ext(filename)
	fileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)
	img, error := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if error != nil {
		return nil, uploadmodel.ErrCannotSaveFile(error)
	}
	img.Width = w
	img.Height = h
	img.Extension = fileExt
	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, error := image.DecodeConfig(reader)
	if error != nil {
		log.Print("err: ", error)
		return 0, 0, error
	}
	return img.Width, img.Height, nil
}
