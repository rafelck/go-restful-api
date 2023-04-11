package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"rafelck/go-restful-api/app"
	"rafelck/go-restful-api/controller"
	"rafelck/go-restful-api/helper"
	"rafelck/go-restful-api/middleware"
	"rafelck/go-restful-api/repository"
	"rafelck/go-restful-api/service"
)

func main() {

	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
