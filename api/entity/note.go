package entity

import "github.com/matthieurobert/amp/api/config"

type Note struct {
	Id     int64  `json:"id"`
	Body   string `json:"body"`
	UserId int64  `json:"user_id"`
}

func GetNotes(userId int64) ([]Note, error) {
	var notes []Note

	err := config.POSTGRES.DB.Model(&notes).Where("user_id = ?", userId).Select()

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
