package restaurantbiz

import (
	"context"
	"errors"
	"fooddelivery/module/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("restaurant name cannot be blank")
	}
	err := biz.store.Create(ctx, data)

	return err
}
