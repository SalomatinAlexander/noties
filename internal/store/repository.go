package store

import (
	"fmt"
	"log"
	"noties/internal/models"
)

type Repository struct {
	*Store
}

func NewRepository(s *Store) *Repository {
	if err := s.Open(); err != nil {
		log.Fatal("Open db error:" + fmt.Sprint(err))
	}
	return &Repository{
		Store: s}
}

func (r *Repository) CreateNewNote(note models.NoteCreate) (int, error) {
	db, err := r.Store.db.Begin()
	if err != nil {
		return -1, err
	}
	createNoteQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, list_id, title, description_text, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		"notes_table")
	var id int
	result := db.QueryRow(createNoteQuery, note.UserId, note.ListId, note.Title,
		note.Description, note.CreateAt, note.UpdateAt)
	if err := result.Scan(&id); err != nil {
		db.Rollback()
		return 0, err
	}
	db.Commit()
	return id, nil
}

func (r *Repository) GetAllNotes() ([]models.Note, error) {
	db, err := r.Store.db.Begin()
	emptyList := make([]models.Note, 0)
	if err != nil {
		return emptyList, err
	}

	getAllQuery := "SELECT * FROM notes_table"
	var list []models.Note
	result := db.QueryRow(getAllQuery)

	if err := result.Scan(&list); err != nil {
		db.Rollback()
		return emptyList, err
	}
	db.Commit()
	return list, nil
}
