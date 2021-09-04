package uploadstorage

import (
	"context"
	"fooddelivery/common"
)

func (s *sqlStore) ListImages(
	ctx context.Context,
	ids []int,
	moreKeys ...string,
) (common.Images, error) {
	db := s.db

	var result common.Images

	db = db.Table(common.Images{}.TableName())

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil

}
