package db

import (
	"emotionDiary/controllers/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func ConnectToDatabase() {

	dsn := "host=localhost user=admin password=admin dbname=postgres port=5432"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Note{}, &models.Sentence{}); err != nil {
		log.Fatal(err.Error())
	}
}

func CreateNote(note *models.Note) error {
	result := db.Create(note)

	return result.Error
}

func GetNote(id *uuid.UUID) (*models.Note, error) {

	var note models.Note

	result := db.First(&note, id)

	return &note, result.Error
}

func GetFirstUser() (*models.User, error) {
	var user models.User
	result := db.First(&user)

	return &user, result.Error
}

func DeleteNote(id *uuid.UUID) error {
	result := db.Delete(&models.Note{}, id)

	if result.RowsAffected == 0 {
		log.Println(result.Error)
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func getSentences(noteId *uuid.UUID, sentences *[]models.Sentence) error {
	result := db.Where("note_id = ?", noteId).Find(sentences)

	return result.Error
}

func DeleteSentences(noteId *uuid.UUID) error {
	result := db.Where("note_id = ?", noteId).Delete(&models.Sentence{})

	return result.Error
}

func UpdateNote(note *models.Note) error {

	result := db.Save(note)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func GetNotes(startTime, endTime *time.Time, notes *[]models.Note) (err error) {
	result := db.Where("created_at BETWEEN ? AND ?", startTime, endTime).Find(notes)

	err = result.Error

	if err != nil {
		return
	}

	for i := range *notes {
		if err = getSentences(&(*notes)[i].ID, &(*notes)[i].Sentences); err != nil {
			return
		}
	}

	return
}
