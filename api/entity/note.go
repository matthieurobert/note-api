package entity

import "github.com/matthieurobert/amp/api/config"

type Note struct {
	Id   int64  `json:"id"`
	Body string `json:"body"`
}

func GetNotes() ([]Note, error) {
	var notes []Note

	err := config.POSTGRES.DB.Model(&notes).Select()

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func PostNote(note *Note) (int64, error) {
	_, err := config.POSTGRES.DB.Model(note).Insert()

	if err != nil {
		return 0, err
	}

	return note.Id, nil
}
