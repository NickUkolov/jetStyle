package controllers

import (
	"net/http"
	"notes-service/models"
	"notes-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	Service *services.NoteService
}

func (ctrl *NoteController) GetAllNotes(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
		return
	}

	var tokenUserID *uint
	if id, exists := c.Get("id"); exists {
		uid := id.(uint)
		tokenUserID = &uid
	}

	query := c.Query("query")
	var queryPtr *string
	if query != "" {
		queryPtr = &query
	}

	queryUserID := c.Query("user_id")

	var queryUserIDPtr *int
	if queryUserID != "" {
		queryUserIDa, err := strconv.Atoi(queryUserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id value"})
			return
		}
		queryUserIDPtr = &queryUserIDa
	}

	notes, err := ctrl.Service.GetAllNotes(limit, offset, tokenUserID, queryUserIDPtr, queryPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func (ctrl *NoteController) GetNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	note, err := ctrl.Service.GetNoteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userID uint
	if id, exists := c.Get("id"); exists {
		uid := id.(uint)
		userID = uid
	}

	if userID != note.AuthorID && note.Public == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (ctrl *NoteController) CreateNote(c *gin.Context) {
	var noteInput models.NoteInput
	if err := c.ShouldBindJSON(&noteInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorID, _ := c.Get("id")

	note := models.Note{
		AuthorID: authorID.(uint),
		Content:  noteInput.Content,
		Public:   noteInput.Public,
	}

	note, err := ctrl.Service.CreateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, note)
}

func (ctrl *NoteController) UpdateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var noteInput models.NoteInput
	if err := c.ShouldBindJSON(&noteInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := ctrl.Service.GetNoteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var userID uint
	if id, exists := c.Get("id"); exists {
		uid := id.(uint)
		userID = uid
	}

	if userID != note.AuthorID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	note.Content = noteInput.Content
	note.Public = noteInput.Public

	note, err = ctrl.Service.UpdateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, note)
}

func (ctrl *NoteController) DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	note, err := ctrl.Service.GetNoteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userID uint
	if id, exists := c.Get("id"); exists {
		uid := id.(uint)
		userID = uid
	}

	if userID != note.AuthorID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if err := ctrl.Service.DeleteNote(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
