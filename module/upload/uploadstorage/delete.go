package uploadstorage

import (
	"context"
	"fooddelivery/common"
)

func (s *sqlStore) DeleteImage(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(common.Image{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
