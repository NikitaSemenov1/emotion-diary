package main

import (
	"emotionDiary/controllers/db"
	"emotionDiary/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("api/register", handlers.RegisterHandler)
	r.POST("api/login", handlers.LoginHandler)

	{
		authorized := r.Group("/", gin.Logger(), gin.Recovery())

		authorized.POST("api/logout", handlers.LogoutHandler)

		authorized.POST("api/note", handlers.CreateNoteHandler)
		authorized.GET("api/note", handlers.GetNotesByDateRangeHandler)
		authorized.PUT("api/note/:id", handlers.UpdateNoteHandler)
		authorized.DELETE("api/note/:id", handlers.DeleteNote)
	}

	return r
}

func main() {
	r := setupRouter()
	db.ConnectToDatabase()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
