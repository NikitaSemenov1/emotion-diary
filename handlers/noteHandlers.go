package handlers

import (
	"emotionDiary/controllers"
	"emotionDiary/handlers/requestEntities"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func errorResponse(message string) gin.H {
	return gin.H{"error": message}
}

func okResponse(message string) gin.H {
	return gin.H{"message": message}
}

func CreateNoteHandler(c *gin.Context) {
	var createNoteEntity requestEntities.CreateNoteEntity

	if err := c.ShouldBindJSON(&createNoteEntity); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	if err := controllers.CreateNote(&createNoteEntity); err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, okResponse("Ok"))
}

func GetNotesByDateRangeHandler(c *gin.Context) {
	var getNotesEntityByDateRangeEntity requestEntities.GetNotesByDateRangeEntity

	if err := c.ShouldBind(&getNotesEntityByDateRangeEntity); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	notes, err := controllers.GetNotesByDateRange(&getNotesEntityByDateRangeEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, notes)
}

func UpdateNoteHandler(c *gin.Context) {
	var updateNoteEntity requestEntities.UpdateNoteEntity

	if err := c.ShouldBind(&updateNoteEntity); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	updateNoteEntity.Id = c.Param("id")

	err := controllers.UpdateNote(&updateNoteEntity)

	if errors.Is(err, controllers.AccessError) {
		c.JSON(http.StatusForbidden, errorResponse(err.Error()))
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, errorResponse(err.Error()))
		return
	}

	if errors.Is(err, controllers.InvalidUUID) {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, okResponse("Updated"))
}

func DeleteNote(c *gin.Context) {

	err := controllers.DeleteNote(c.Param("id"))

	if errors.Is(err, controllers.AccessError) {
		c.JSON(http.StatusForbidden, errorResponse(err.Error()))
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, errorResponse(err.Error()))
		return
	}

	if errors.Is(err, controllers.InvalidUUID) {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, okResponse("Ok"))
}
