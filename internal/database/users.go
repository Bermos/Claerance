package database

import (
	"Claerance/internal/users"
	"database/sql"
	"errors"
	"fmt"
)

func (d database) AddUser(username string, hashedPW string) error {
	user, err := d.GetUserByName(username)

	if err == nil && user.PwdHash != "" {
		return errors.New("username already exists")
	}

	_, err = d.db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, username, hashedPW)
	return err
}

func (d database) GetUserByName(username string) (users.User, error) {
	query := `SELECT * FROM users WHERE username = ?`
	row := d.db.QueryRow(query, username)

	return getUserFromQueryRow(row)
}

func (d database) GetUserById(userId int) (users.User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	row := d.db.QueryRow(query, userId)

	return getUserFromQueryRow(row)
}

func (d database) GetAllUsers() ([]users.User, error) {
	query := `SELECT * FROM users`
	rows, err := d.db.Query(query)

	var us []users.User
	for rows.Next() {
		user, _ := getUserFromQueryRows(rows)
		us = append(us, user)
	}

	return us, err
}

func (d database) UpdateUser(user users.User) error {
	query := `UPDATE users SET email = ?, telegram_id = ? WHERE id = ?`
	result, err := d.db.Exec(query, user.Id)

	if rowsAffected, _ := result.RowsAffected(); err == nil && rowsAffected != 1 {
		err = fmt.Errorf("no changes made")
	}

	return err
}

func (d database) DeleteUserById(userId int) bool {
	query := `DELETE FROM users WHERE id = ?`
	result, _ := d.db.Exec(query, userId)
	rowsAffected, err := result.RowsAffected()

	return rowsAffected >= 1 && err == nil
}

func getUserFromQueryRow(row *sql.Row) (users.User, error) {
	var u users.User
	var email sql.NullString
	var telegramID sql.NullInt32

	err := row.Scan(&u.Id, &u.Username, &u.PwdHash, &u.CreatedAt, &email, &telegramID)

	if email.Valid {
		u.Email = email.String
	}

	if telegramID.Valid {
		u.TelegramId = int(telegramID.Int32)
	}

	return u, err
}

func getUserFromQueryRows(row *sql.Rows) (users.User, error) {
	var u users.User
	var email sql.NullString
	var telegramID sql.NullInt32

	err := row.Scan(&u.Id, &u.Username, &u.PwdHash, &u.CreatedAt, &email, &telegramID)

	if email.Valid {
		u.Email = email.String
	}

	if telegramID.Valid {
		u.TelegramId = int(telegramID.Int32)
	}

	return u, err
}
