package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/controller"
)

func ApiRouter(app *gin.Engine) {
	api := app.Group("/api")
	{
		job := api.Group("/job")
		{
			// Create a new job
			job.POST("", controller.CreateJob)
			//Get the job list
			job.GET("", controller.GetJobList)
			//Get the job by id
			job.GET("/:id", controller.GetJobById)
			//Update the job
			job.PATCH("/:id", controller.UpdateJobById)
			//Delete the job
			job.DELETE("/:id", controller.DeleteJobById)
		}
	}
}
