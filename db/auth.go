package db

import (
	"database/sql"
	"errors"

	"github.com/sayaandreas/goingnowhere/models"
)

func (db Database) GetUserByUsername(username string) (models.User, error) {
	u := models.User{}
	query := `SELECT * FROM users WHERE username = $1;`
	row := db.Conn.QueryRow(query, username)
	switch err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt); err {
	case sql.ErrNoRows:
		return u, errors.New("No match row")
	default:
		return u, err
	}
}
