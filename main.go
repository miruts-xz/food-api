package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/comment/repository"
	"github.com/miruts/food-api/comment/service"
	"github.com/miruts/food-api/delivery/http/handler"
	repository2 "github.com/miruts/food-api/user/repository"
	service2 "github.com/miruts/food-api/user/service"
	"net/http"
	"os"
)

const (
	host     = "ijobs.cx20b90aqxzy.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "ijobspassword"
	dbname   = "food_api"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbconn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	//dbconn.CreateTable(&entity.Category{}, &entity.Ingredient{}, &entity.Item{}, &entity.Order{}, &entity.Role{}, &entity.Error{}, &entity.User{}, &entity.Comment{})

	roleRepo := repository2.NewRoleGormRepo(dbconn)
	roleSrv := service2.NewRoleService(roleRepo)
	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)

	commentRepo := repository.NewCommentGormRepo(dbconn)
	commentSrv := service.NewCommentService(commentRepo)
	adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)

	router := httprouter.New()

	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)

	router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	router.POST("/v1/admin/comments", adminCommentHandler.PostComment)
	router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)
	router.GET("/favicon.ico", adminCommentHandler.GetComments)
	router.GET("/", adminCommentHandler.GetComments)

	http.ListenAndServe(":"+port, router)
}
