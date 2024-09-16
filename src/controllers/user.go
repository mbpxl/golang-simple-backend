package controllers

import (
	"net/http"

	"main/src/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	loggedUser, err := input.Login()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": loggedUser})
}

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	loggedUser, err := input.Register()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": loggedUser})
}

func GetProfile(c *gin.Context) {
	// userId := 1
	// // if userId == nil {
	// //  c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
	// //  return
	// // }
	// startup, err := models.FetchUser(userId)
	// if err != nil {
	//  c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	//  return
	// }
	// c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "status": "success", "data": startup})
}
