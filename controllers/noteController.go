package controllers

import (
	"emotionDiary/controllers/db"
	"emotionDiary/controllers/emotion"
	"emotionDiary/controllers/models"
	"emotionDiary/handlers/requestEntities"
	"emotionDiary/ml"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

func parseSentences(text string, noteId uuid.UUID) []models.Sentence {

	// TODO: parse note into sentences

	sentenceEmotion, ok := emotion.EmotionMapper[ml.GetEmotion(&text)]

	if !ok {
		log.Println("Got an unknown emotion. Set to default")
		sentenceEmotion = emotion.NoEmotion
	}

	return []models.Sentence{
		{
			Index:   0,
			Text:    text,
			Emotion: sentenceEmotion,
			NoteId:  noteId,
		},
	}
}

func CreateNote(entity *requestEntities.CreateNoteEntity) error {
	user, _ := getUser()

	note := models.Note{
		UserId: user.ID,
	}

	note.Sentences = parseSentences(entity.Text, note.ID)
	note.Emotion = emotion.GetNoteEmotion(&note)

	err := db.CreateNote(&note)

	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateNote(entity *requestEntities.UpdateNoteEntity) error {
	user, _ := getUser()

	id, err := uuid.Parse(entity.Id)

	if errors.Is(err, InvalidUUID) {
		log.Println(err)
		return err
	}

	note, err := db.GetNote(&id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		return err
	}

	if note.UserId != user.ID {
		log.Println("Access error")
		return AccessError
	}

	if err := db.DeleteSentences(&note.ID); err != nil {
		log.Println(err)
		return err
	}

	note.Sentences = parseSentences(entity.Text, user.ID)

	err = db.UpdateNote(note)

	if err != nil {
		log.Println(err)
	}

	return err
}

func DeleteNote(idStr string) error {
	id, err := uuid.Parse(idStr)

	if err != nil {
		log.Println(err)
		return InvalidUUID
	}

	err = db.DeleteNote(&id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetNotesByDateRange(entity *requestEntities.GetNotesByDateRangeEntity) (interface{}, error) {
	var notes []models.Note

	err := db.GetNotes(&entity.DateStart, &entity.DateEnd, &notes)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return notes, nil
}
