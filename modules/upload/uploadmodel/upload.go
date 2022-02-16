package uploadmodel

import (
	"errors"
	"hfs_backend/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (receiver Upload) TableName() string {
	return "uploads"
}

//func (receiver *Upload) Mask(isAdmin bool)  {
//  receiver.GenUID(common.ErrDB())
//}
var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge")
)

func ErrFileIsnotImage(err error) *common.AppError {
	return common.NewCustomError(err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}
func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err,
		"cannot save upload file",
		"ErrCannotSaveFile",
	)
}
