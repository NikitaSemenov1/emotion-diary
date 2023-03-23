package controllers

import (
	"emotionDiary/controllers/db"
	"emotionDiary/controllers/models"
)

func getUser() (*models.User, error) {
	return db.GetFirstUser()
}
