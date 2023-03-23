package emotion

import "emotionDiary/controllers/models"

const (
	NoEmotion = iota
	Joy
	Sadness
)

type Type = uint

var EmotionMapper = map[string]Type{
	"no_emotion": NoEmotion,
	"sadness":    Sadness,
	"Joy":        Joy,
}

func GetNoteEmotion(note *models.Note) Type {
	return NoEmotion
}
