package restaurantlikestorage

import (
	"context"
	"fmt"
	"fooddelivery/common"
	"fooddelivery/module/restaurantlike/restaurantlikemodel"

	"github.com/btcsuite/btcutil/base58"
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

func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.Like

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB((err))
	}

	// for i := range moreKeys {
	// 	db = db.Preload(moreKeys[i])
	// }

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("created_at < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		users[i] = *result[i].User
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format("2006-01-02T15:04:05.999999"))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
