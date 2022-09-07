package controllers

import (
	"encoding/json"
	"example.com/crud-with-mongo/models"
	// "fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// METHODs

func (uc UserController) GetUser(c *gin.Context) {
	// get parameter
	id := c.Param("id")

	// validate id
	if !bson.IsObjectIdHex(id) {
		c.IndentedJSON(404, gin.H{"message": "invalid id parameter"})
		return
	}

	oid := bson.ObjectIdHex(id)

	user_model := models.User{}

	err := uc.Session.DB("crud-with-golang").C("users").FindId(oid).One(&user_model)
	if err != nil {
		c.IndentedJSON(404, gin.H{"message": err})
		return
	}

	uj, err := json.Marshal(user_model)
	if err != nil {
		c.IndentedJSON(404, gin.H{"message": "error"})
		return
	}

	// success
	c.IndentedJSON(http.StatusOK, gin.H{"user": uj})
}

func (uc UserController) CreateUser(c *gin.Context) {
	var new_user models.User
	new_user.Id = bson.NewObjectId() // make id automatically with bson

	// make new User struct from incoming data automatically
	err := c.BindJSON(&new_user) // return error information

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	// insert into database
	err = uc.Session.DB("crud-with-golang").C("users").Insert(new_user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error while inserting new user"})
		return
	}

	c.IndentedJSON(http.StatusOK, new_user)
}

func (uc UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// validate id
	if !bson.IsObjectIdHex(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid id parameter"})
		return
	}

	oid := bson.ObjectIdHex(id)

	err := uc.Session.DB("crud-with-golang").C("users").RemoveId(oid)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to remove user"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "failed to remove user"})
}
