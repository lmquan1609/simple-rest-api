package restaurantbiz

import "context"

type DeleteRestaurantStore interface {
	SoftDeleteData(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	err := biz.store.SoftDeleteData(ctx, id)
	return err
}
