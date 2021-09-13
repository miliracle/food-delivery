package restaurantlikestorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) GetRestaurantsLike(ctx context.Context, ids []int) (map[int]int, error) {
	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName())

	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"colume:count;"`
	}

	var restaurantLikes []sqlData

	err := db.Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&restaurantLikes).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	for _, r := range restaurantLikes {
		result[r.RestaurantId] = r.LikeCount
	}

	return result, nil
}
