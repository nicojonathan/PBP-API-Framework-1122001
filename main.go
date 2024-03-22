package main

import (
	"belajar_gin/controllers"
	"log"

	"github.com/gin-gonic/gin"
	//"github.com/gorilla/mux"
)

func main() {
	//router := mux.NewRouter()
	router := gin.Default()

	// READ
	// router.GET("/user", controllers.GetAllUsers)

	// CREATE
	// router.POST("/user", controllers.InsertUser)

	// UPDATE
	// router.PUT("/user", controllers.UpdateUser)

	// DELETE
	// router.DELETE("/user", controllers.DeleteUser)


	// Explore Fitur Routes Grouping
	rGroup := router.Group("/user")

	// READ
	rGroup.GET("", controllers.GetAllUsers)

	// CREATE
	rGroup.POST("", controllers.InsertUser)

	// UPDATE
	rGroup.PUT("", controllers.UpdateUser)

	// DELETE
	rGroup.DELETE("", controllers.DeleteUser)

	err := router.Run("localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
}
