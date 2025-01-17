package categories

import (
	"context"
	"math"
	"social-network/internal/model/categories"
	"social-network/internal/model/paginate"

	"github.com/rs/zerolog/log"
)

func (s *service) ListCategory(ctx context.Context, limit, page int) (*categories.ListCategoriesRes, error) {
	offset := (page - 1) * limit
	listCategories, err := s.categoryRepo.ListCategory(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, nil

	}
	total, err := s.categoryRepo.TotalCategory(ctx)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	totalPages := math.Ceil(float64(total) / float64(limit))

	result := categories.ListCategoriesRes{
		Paginate: paginate.Paginate{
			Limit:       limit,
			TotalPage:   int(totalPages),
			CurrentPage: page,
		},
		Data: *listCategories,
	}

	return &result, nil
}
