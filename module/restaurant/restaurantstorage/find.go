package restaurantstorage

import (
	"context"
	"fooddelivery/module/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	var result *restaurantmodel.Restaurant

	if err := db.First(&result).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {

		// }
		return nil, err
	}

	return result, nil
}
