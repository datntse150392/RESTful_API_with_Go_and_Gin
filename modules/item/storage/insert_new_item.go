package storage

import (
	"RESTful_API_with_Go_and_Gin/modules/item/model"
	"context"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
