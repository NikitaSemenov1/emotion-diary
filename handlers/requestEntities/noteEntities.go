package requestEntities

import "time"

type CreateNoteEntity struct {
	Text string `binding:"required"`
}

type UpdateNoteEntity struct {
	CreateNoteEntity
	Id string
}

type GetNotesByDateRangeEntity struct {
	DateStart time.Time `form:"dateStart" binding:"required"`
	DateEnd   time.Time `form:"dateEnd" binding:"required"`
}
