package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/menu"
)

// ItemGormRepo implements the menu.ItemRepository interface
type ItemGormRepo struct {
	conn *gorm.DB
}

// NewItemGormRepo will create a new object of ItemGormRepo
func NewItemGormRepo(db *gorm.DB) menu.ItemRepository {
	return &ItemGormRepo{conn: db}
}

// Items returns all food menus stored in the database
func (itemRepo *ItemGormRepo) Items() ([]entity.Item, []error) {
	items := []entity.Item{}
	errs := itemRepo.conn.Find(&items).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	for j, i := range items {
		categs := []entity.Category{}
		ings := []entity.Ingredient{}
		_ = itemRepo.conn.Model(&i).Related(&categs, "Categories")
		_ = itemRepo.conn.Model(&i).Related(&ings, "Ingredients")
		items[j].Categories = categs
		items[j].Ingredients = ings
	}
	return items, errs
}

// Item retrieves a food menu by its id from the database
func (itemRepo *ItemGormRepo) Item(id uint) (*entity.Item, []error) {
	item := entity.Item{}
	ings := []entity.Ingredient{}
	categs := []entity.Category{}
	errs := itemRepo.conn.First(&item, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	_ = itemRepo.conn.Model(&item).Related(&categs,"Categories")
	_ = itemRepo.conn.Model(&item).Related(&ings, "Ingredients")
	item.Ingredients = ings
	item.Categories = categs
	return &item, errs
}

// UpdateItem updates a given food menu item in the database
func (itemRepo *ItemGormRepo) UpdateItem(item *entity.Item) (*entity.Item, []error) {
	itm := item
	errs := itemRepo.conn.Save(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// DeleteItem deletes a given food menu item from the database
func (itemRepo *ItemGormRepo) DeleteItem(id uint) (*entity.Item, []error) {
	itm, errs := itemRepo.Item(id)

	if len(errs) > 0 {
		return nil, errs
	}
	errs = itemRepo.conn.Delete(itm, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// StoreItem stores a given food menu item in the database
func (itemRepo *ItemGormRepo) StoreItem(item *entity.Item) (*entity.Item, []error) {
	itm := item
	errs := itemRepo.conn.Create(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
