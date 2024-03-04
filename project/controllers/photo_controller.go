package controllers

import (
	"net/http"
	"project/database"
	"project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {
	var newPhoto models.Photo
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	path := "uploads/" + file.Filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newPhoto.PhotoUrl = path
	newPhoto.Title = c.PostForm("title")
	newPhoto.Caption = c.PostForm("caption")
	newPhoto.UserID = 1

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo info to database."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Photo uploaded successfully", "photo": newPhoto})
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	id, err := strconv.Atoi(photoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	if err := database.DB.Delete(&models.Photo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully", "photoId": photoID})
}
