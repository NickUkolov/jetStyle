package repositories

import (
	"gorm.io/gorm"
	"notes-service/models"
)

type NoteRepository struct {
	DB *gorm.DB
}

func (repo *NoteRepository) GetAllNotes(limit, offset int, tokenUserID *uint, queryUserID *int, query *string) ([]models.Note, error) {
	var notes []models.Note
	db := repo.DB

	if tokenUserID != nil {
		db = db.Where("public = ? OR (public = false AND author_id = ?)", true, *tokenUserID)
	} else {
		db = db.Where("public = ?", true)
	}

	if query != nil {
		db = db.Where("content ILIKE ?", "%"+*query+"%")
	}

	if queryUserID != nil {
		db = db.Where("author_id = ?", *queryUserID)
	}

	result := db.Limit(limit).Offset(offset).Find(&notes)
	return notes, result.Error
}

func (repo *NoteRepository) GetNoteByID(id uint) (models.Note, error) {
	var note models.Note
	result := repo.DB.First(&note, id)
	return note, result.Error
}

func (repo *NoteRepository) CreateNote(note models.Note) (models.Note, error) {
	result := repo.DB.Create(&note)
	return note, result.Error
}

func (repo *NoteRepository) UpdateNote(note models.Note) (models.Note, error) {
	result := repo.DB.Save(&note)
	return note, result.Error
}

func (repo *NoteRepository) DeleteNote(id uint) error {
	result := repo.DB.Delete(&models.Note{}, id)
	return result.Error
}
