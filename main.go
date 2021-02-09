package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/comment/repository"
	"github.com/miruts/food-api/comment/service"
	"github.com/miruts/food-api/delivery/http/handler"
	repository3 "github.com/miruts/food-api/menu/repository"
	service3 "github.com/miruts/food-api/menu/service"
	repository4 "github.com/miruts/food-api/order/repository"
	service4 "github.com/miruts/food-api/order/usecase"
	repository2 "github.com/miruts/food-api/user/repository"
	service2 "github.com/miruts/food-api/user/service"
	"net/http"
	"os"
)

const (
	host     = "foodorder.cx20b90aqxzy.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "foodorder"
	dbname   = "restaurantdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbconn, err := gorm.Open("postgres", psqlInfo)
	//dbconn, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost/restaurantdb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	commentRepo := repository.NewCommentGormRepo(dbconn)
	menuRepo := repository3.NewCategoryGormRepo(dbconn)
	userRepo := repository2.NewUserGormRepo(dbconn)
	orderRepo := repository4.NewOrderGormRepo(dbconn)
	ingredientRepo := repository3.NewIngredientGormRepo(dbconn)
	itemRepo := repository3.NewItemGormRepo(dbconn)
	roleRepo := repository2.NewRoleGormRepo(dbconn)

	commentSrv := service.NewCommentService(commentRepo)
	menuServ := service3.NewCategoryService(menuRepo)
	userServ := service2.NewUserService(userRepo)
	ingredientServ := service3.NewIngredientService(ingredientRepo)
	itemServ := service3.NewItemService(itemRepo)
	orderServ := service4.NewOrderService(orderRepo)
	roleSrv := service2.NewRoleService(roleRepo)

	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)
	adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)
	adminMenuHandler := handler.NewAdminMenuHandler(menuServ)
	adminOrderHandler := handler.NewAdminOrdertHandler(orderServ)
	adminItemHandler := handler.NewAdminItemHandler(itemServ)
	adminIngredientHandler := handler.NewAdminIngredientHandler(ingredientServ)
	adminUserHandler := handler.NewAdminUserHandler(userServ)

	router := httprouter.New()

	router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	router.POST("/v1/admin/comments", adminCommentHandler.PostComment)
	router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)

	router.GET("/v1/admin/roles/:id", adminRoleHandler.GetSingleRole)
	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)
	router.PUT("/v1/admin/roles/:id", adminRoleHandler.PutRole)
	router.POST("/v1/admin/roles", adminRoleHandler.PostRole)
	router.DELETE("/v1/admin/roles/:id", adminRoleHandler.DeleteRole)

	router.GET("/v1/admin/categories/:id", adminMenuHandler.GetSingleMenu)
	router.GET("/v1/admin/categories", adminMenuHandler.GetMenus)
	router.PUT("/v1/admin/categories/:id", adminMenuHandler.PutMenu)
	router.POST("/v1/admin/categories", adminMenuHandler.PostMenu)
	router.DELETE("/v1/admin/categories/:id", adminMenuHandler.DeleteMenu)

	router.GET("/v1/admin/orders/:id", adminOrderHandler.GetSingleOrder)
	router.GET("/v1/admin/orders", adminOrderHandler.GetOrders)
	router.PUT("/v1/admin/orders/:id", adminOrderHandler.PutOrder)
	router.POST("/v1/admin/orders", adminOrderHandler.PostOrder)
	router.DELETE("/v1/admin/orders/:id", adminOrderHandler.DeleteOrder)

	router.GET("/v1/admin/items/:id", adminItemHandler.GetSingleItem)
	router.GET("/v1/admin/items", adminItemHandler.GetItems)
	router.PUT("/v1/admin/items/:id", adminItemHandler.PutItems)
	router.POST("/v1/admin/items", adminItemHandler.PostItem)
	router.DELETE("/v1/admin/items/:id", adminItemHandler.DeleteItem)

	router.GET("/v1/admin/ingredients/:id", adminIngredientHandler.GetSingleIngredient)
	router.GET("/v1/admin/ingredients", adminIngredientHandler.GetIngredients)
	router.PUT("/v1/admin/ingredients/:id", adminIngredientHandler.PutIIngredients)
	router.POST("/v1/admin/ingredients", adminIngredientHandler.PostIngredient)
	router.DELETE("/v1/admin/ingredients/:id", adminIngredientHandler.DeleteIngredient)

	router.GET("/v1/admin/username/:username", adminUserHandler.GetByUsername)
	router.GET("/v1/admin/users/:id", adminUserHandler.GetSingleUser)
	router.GET("/v1/admin/users", adminUserHandler.GetUsers)
	router.PUT("/v1/admin/users/:id", adminUserHandler.PutUser)
	router.POST("/v1/admin/users", adminUserHandler.PostUser)
	router.DELETE("/v1/admin/users/:id", adminUserHandler.DeleteUser)

	router.GET("/", handler.Index)

	http.ListenAndServe(":"+port, router)
}
