package routes

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"notes-service/controllers"
	"notes-service/middlewares"
	"notes-service/repositories"
	"notes-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	noteRepo := &repositories.NoteRepository{DB: db}
	noteService := &services.NoteService{Repo: noteRepo}
	noteController := &controllers.NoteController{Service: noteService}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/notes", middlewares.OptionalJWTAuthMiddleware(), noteController.GetAllNotes)
	router.GET("/notes/:id", middlewares.OptionalJWTAuthMiddleware(), noteController.GetNoteByID)

	authRoutes := router.Group("/")
	authRoutes.Use(middlewares.JWTAuthMiddleware())
	{
		authRoutes.POST("/notes", noteController.CreateNote)
		authRoutes.PUT("/notes/:id", noteController.UpdateNote)
		authRoutes.DELETE("/notes/:id", noteController.DeleteNote)
	}
}
