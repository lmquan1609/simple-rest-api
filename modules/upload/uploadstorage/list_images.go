package uploadstorage

import (
	"context"
	"simple-rest-api/common"
)

func (s *sqlStore) ListImages(ctx context.Context, ids []int, moreKeys ...string) ([]common.Image, error) {
	db := s.db
	var result []common.Image

	db = db.Table(common.Image{}.TableName())

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
