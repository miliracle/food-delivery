package restaurantbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurant/restaurantmodel"
	"log"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantsLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)

	}

	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantsLike(ctx, ids)

	if err != nil {
		log.Println("Cannot get restaurants like:", err)
	}

	if mapResLike != nil {
		for i, item := range result {
			result[i].LikedCount = mapResLike[item.Id]
		}
	}

	return result, nil
}
