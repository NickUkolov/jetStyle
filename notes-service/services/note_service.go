package services

import (
	"notes-service/models"
	"notes-service/repositories"
)

type NoteService struct {
	Repo *repositories.NoteRepository
}

func (s *NoteService) GetAllNotes(limit, offset int, userID *uint, queryUserID *int, query *string) ([]models.Note, error) {
	return s.Repo.GetAllNotes(limit, offset, userID, queryUserID, query)
}

func (s *NoteService) GetNoteByID(id uint) (models.Note, error) {
	return s.Repo.GetNoteByID(id)
}

func (s *NoteService) CreateNote(note models.Note) (models.Note, error) {
	return s.Repo.CreateNote(note)
}

func (s *NoteService) UpdateNote(note models.Note) (models.Note, error) {
	return s.Repo.UpdateNote(note)
}

func (s *NoteService) DeleteNote(id uint) error {
	return s.Repo.DeleteNote(id)
}
