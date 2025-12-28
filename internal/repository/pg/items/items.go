package items

import (
	"github.com/Archiker-715/rest-api/internal/entity"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) GetItems() (items []entity.Item, err error) {
	if err = r.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return
}

func (r *ItemRepository) GetByID(id uint) (item *entity.Item, err error) {
	if err = r.DB.Find(&item, id).Error; err != nil {
		return nil, err
	}
	return
}

func (r *ItemRepository) Create(item *entity.Item) error {
	return r.DB.Create(item).Error
}

func (r *ItemRepository) Update(item *entity.Item) error {
	return r.DB.Save(item).Error
}

func (r *ItemRepository) Delete(id uint) error {
	return r.DB.Delete(&entity.Item{ID: id}).Error
}
