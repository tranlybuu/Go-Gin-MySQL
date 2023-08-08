package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func GetCourseList(c *gin.Context) {
	results := model.FindCourseList()
	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func GetCourseById(c *gin.Context) {
	id := c.Param("id")
	result := model.FindCourseDetail(id)
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func CreateCourse(c *gin.Context) {
	var course model.Course
	if err := c.ShouldBind(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	course.ID = primitive.NewObjectID()
	course.Deleted = false
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	message := model.CreateCourse(course)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func DeleteCourseById(c *gin.Context) {
	id := c.Param("id")
	message := model.DeleteCourseById(id)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func UpdateCourseById(c *gin.Context) {
	id := c.Param("id")
	var course model.Course
	if err := c.ShouldBind(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	course.ID, _ = primitive.ObjectIDFromHex(id)
	course.UpdatedAt = time.Now()
	message := model.UpdateCourseById(course)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
