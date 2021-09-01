package restaurantbiz

import (
	"context"
	"fooddelivery/common"
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
	if err := data.CreateValidate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	err := biz.store.Create(ctx, data)

	return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
}
