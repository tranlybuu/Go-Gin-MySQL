package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-gin/initializer"
	"go-gin/model"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var db *gorm.DB

func init() {
	db = initializer.ConnectDb()
}

func GetJobList(c *gin.Context) {
	var result []model.Job

	if err := db.Order("created_at desc").Find(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func CreateJob(c *gin.Context) {
	var data model.JobCreation

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data.ID = uuid.New().String()
	data.Status = false
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	if err := db.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Created successfully",
	})
}

func GetJobById(c *gin.Context) {
	var data model.JobCreation

	id := c.Param("id")

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func UpdateJobById(c *gin.Context) {
	var data model.JobUpdate

	id := c.Param("id")
	data.UpdatedAt = time.Now()
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Updated successfully",
	})
}

func DeleteJobById(c *gin.Context) {
	id := c.Param("id")

	if err := db.Table(model.Job{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Deleted successfully",
	})
}
