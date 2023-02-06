package uploadmodel

import "simple-rest-api/common"

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, "file is not images", "ErrFileIsNotImage")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "cannot save upload file", "ErrCannotSaveFile")
}
