package main

import (
	"belajar_gin/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//"github.com/gorilla/mux"
)

func main() {
	//router := mux.NewRouter()
	router := gin.Default()

	router.Use(RecoveryHandler())
	
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

	// Test for panic
	router.GET("/panic", func(c *gin.Context) {
		panic("Oops! Something went wrong!")
	})


	err := router.Run("localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
}


func RecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				fmt.Fprintf(os.Stderr, "Panic: %v\n", err)
				// Abort the request and return a 500 Internal Server Error response
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		// Continue to the next handler
		c.Next()
	}
}

// func RecoveryHandler(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				// Log the error
// 				fmt.Fprintf(os.Stderr, "Panic: %v\n", err)
// 				// Return a 500 Internal Server Error response
// 				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			}
// 		}()

// 		// Call the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }
