package userstorage

import (
	"Ecommerce/commons"
	usermodel "Ecommerce/module/model"
	"context"
)

func (s *sqlStore) Updates(ctx context.Context, data *usermodel.UserUpdate, id int) (*usermodel.UserUpdate, error) {
	if err := s.db.Where("id=?", id).Updates(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return data, nil
}
