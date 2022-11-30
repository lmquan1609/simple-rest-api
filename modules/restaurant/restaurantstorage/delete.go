package restaurantstorage

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) SoftDeleteData(ctx context.Context, id int) error {
	db := s.db

	err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error

	if err != nil {
		return err
	}
	return nil
}
