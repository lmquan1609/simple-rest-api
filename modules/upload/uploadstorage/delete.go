package uploadstorage

import (
	"context"
	"simple-rest-api/common"
)

func (s *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := s.db

	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
