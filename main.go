package main

import (
	"example.com/crud-with-mongo/controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	// "net/http"
	// "fmt"
)

func getSession() *mgo.Session {
	// connect to mongodb
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	router := gin.Default()

	uc := controllers.NewUserController(getSession())

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	router.Run("localhost:7000")
}
