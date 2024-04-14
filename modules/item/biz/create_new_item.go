package biz

import (
	"RESTful_API_with_Go_and_Gin/modules/item/model"
	"context"
	"errors"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateTodoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if data.Title == "" {
		return errors.New("Title can not be blank")
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
