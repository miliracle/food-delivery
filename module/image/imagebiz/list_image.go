package imagebiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurant/restaurantmodel"
)

type ListImageStore interface {
	ListImages(
		ctx context.Context,
		ids *[]int,
		moreKeys ...string,
	) (common.Images, error)
}

type listImageBiz struct {
	store ListImageStore
}

func NewListImageBiz(store ListImageStore) *listImageBiz {
	return &listImageBiz{store: store}
}

func (biz *listImageBiz) ListImage(
	ctx context.Context,
	ids *[]int,
) (common.Images, error) {
	result, err := biz.store.ListImages(ctx, ids)

	if err != nil {
		return nil, common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)

	}

	return result, nil
}
