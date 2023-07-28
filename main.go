package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	ID        string    `json:"id"  gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Status    bool      `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (TodoItem) TableName() string { return "job" }

type TodoItemCreation struct {
	ID        string    `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Status    bool      `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Name      *string   `json:"name" gorm:"name"`
	Status    *bool     `json:"status" gorm:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Pagination) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 5
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dbUrl := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		job := api.Group("/job")
		{
			// Create a new job
			job.POST("", CreateJob(db))
			//Get the job list
			job.GET("", GetJobList(db))
			//Get the job by id
			job.GET("/:id", GetJobById(db))
			//Update the job
			job.PATCH("/:id", UpdateJobById(db))
			//Delete the job
			job.DELETE("/:id", DeleteJobById(db))
		}
	}
	r.Run("localhost:4000")
}

func GetJobList(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var pagi Pagination
		if err := c.ShouldBind(&pagi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		pagi.Process()

		var result []TodoItem

		if err := db.Table(TodoItem{}.TableName()).Count(&pagi.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Order("created_at desc").
			Offset((pagi.Page - 1) * pagi.Limit).
			Limit(pagi.Limit).
			Find(&result).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":       result,
			"pagination": pagi,
		})
	}
}

func CreateJob(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation

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
}

func GetJobById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation

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
}

func UpdateJobById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemUpdate

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
}

func DeleteJobById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": "Deleted successfully",
		})
	}
}
