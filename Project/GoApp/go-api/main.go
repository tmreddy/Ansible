package main

import (	
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"go-api/db"
	"go-api/models"
)	

func main(){
	r := gin.Default()

	r.GET("/users", getUsersHandler)
	r.GET("/users/:id", getUserHandler)
	r.POST("/users", createUserHandler)
	r.PUT("/users/:id", updateUserHandler)
	r.DELETE("/users", deleteUserHandler)

	r.Run(":8000")
}

func getUsersHandler(c *gin.Context) {
	users := db.GetUsers()
	c.JSON(http.StatusOK, users)
}	

func getUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err:= strconv.Atoi(idStr)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}		
	
	user := db.GetUser(id)
	
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createUserHandler(c *gin.Context) {
	var user models.User
	
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user.ID = len(db.GetUsers()) + 1
	db.AddUser(user)

	c.JSON(http.StatusCreated, user)
}	

func updateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	
	var user models.User
	
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	updated := db.UpdateUser(id, user)

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	
	deleted := db.DeleteUser(id)
	
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}	