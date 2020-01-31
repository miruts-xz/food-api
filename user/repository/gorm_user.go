package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/user"
)

// UserGormRepo Implements the menu.UserRepository interface
type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates a new object of UserGormRepo
func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

// Users return all users from the database
func (userRepo *UserGormRepo) Users() ([]entity.User, []error) {
	users := []entity.User{}
	errs := userRepo.conn.Find(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	for i, u := range users {
		roles := []entity.Role{}
		orders := []entity.Order{}
		_ = userRepo.conn.Model(&u).Related(&roles, "Roles")
		_ = userRepo.conn.Model(&u).Related(&orders, "Orders")
		users[i].Roles = roles
		users[i].Orders = orders
	}
	return users, errs
}

// User retrieves a user by its id from the database
func (userRepo *UserGormRepo) User(id uint) (*entity.User, []error) {
	user := entity.User{}
	roles := []entity.Role{}
	orders := []entity.Order{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	_ = userRepo.conn.Model(&user).Related(&roles, "Roles").GetErrors()
	_ = userRepo.conn.Model(&user).Related(&orders, "Orders")
	user.Orders = orders
	user.Roles = roles
	return &user, errs
}

// UpdateUser updates a given user in the database
func (userRepo *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given user from the database
func (userRepo *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = userRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a new user into the database
func (userRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (userRepo *UserGormRepo) UserByUsername(uname string) (*entity.User, []error) {
	usr := &entity.User{}
	errs := userRepo.conn.Where("user_name = ?", uname).First(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
