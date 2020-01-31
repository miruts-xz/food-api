package entity

import "time"

// Category represents Food Menu Category
type Category struct {
	ID          uint `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description"`
	Image       string `json:"image" gorm:"type:varchar(255)"`
	Items       []Item `json:"items" gorm:"many2many:item_categories"`
}

// Role repesents application user roles
type Role struct {
	ID   uint `json:"id"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}

// Item represents food menu items
type Item struct {
	ID          uint `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Price       float32 `json:"price"`
	Description string `json:"description"`
	Categories  []Category   `json:"categories" gorm:"many2many:item_categories"`
	Image       string       `json:"image" gorm:"type:varchar(255)"`
	Ingredients []Ingredient `json:"ingredients" gorm:"many2many:item_ingredients"`
}

// Ingredient represents ingredients in a food item
type Ingredient struct {
	ID          uint `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description"`
}

// Order represents customer order
type Order struct {
	ID       uint `json:"id"`
	PlacedAt time.Time `json:"placed_at"`
	UserID   uint `json:"user_id"`
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

// User represents application user
type User struct {
	ID       uint `json:"id"`
	UserName string `json:"username" gorm:"type:varchar(255);not null"`
	FullName string `json:"fullname" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null; unique"`
	Phone    string `json:"phone" gorm:"type:varchar(100);not null; unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles"`
	Orders   []Order `json:"orders"`
}

// Comment represents comments forwarded by application users
type Comment struct {
	ID       uint      `json:"id"`
	FullName string    `json:"fullname" gorm:"type:varchar(255)"`
	Message  string    `json:"message"`
	Phone    string    `json:"phone" gorm:"type:varchar(100);not null; unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null; unique"`
	PlacedAt time.Time `json:"placedat"`
}

// Error represents error message
type Error struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}
