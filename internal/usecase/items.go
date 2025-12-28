package usecase

import (
	"errors"

	"github.com/Archiker-715/rest-api/internal/entity"
	"github.com/Archiker-715/rest-api/internal/repository/pg/items"
)

type ItemService struct {
	repo *items.ItemRepository
}

func NewItemService(repo *items.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) GetItems() ([]entity.Item, error) {
	return s.repo.GetItems()
}

func (s *ItemService) CreateItem(req *entity.ItemRequest) (*entity.Item, error) {
	if err := s.validateItemRequest(req); err != nil {
		return nil, err
	}
	item := &entity.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
	return item, s.repo.Create(item)
}

func (s *ItemService) UpdateItem(req *entity.ItemRequest, id int) (*entity.Item, error) {
	if err := s.validateItemRequest(req); err != nil {
		return nil, err
	}
	item := &entity.Item{
		ID:          uint(id),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	return item, s.repo.Update(item)
}

func (s *ItemService) DeleteItem(id int) error {
	return s.repo.Delete(uint(id))
}

func (s *ItemService) validateItemRequest(req *entity.ItemRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if req.Price < 0 {
		return errors.New("price must be non-negative")
	}
	return nil
}
