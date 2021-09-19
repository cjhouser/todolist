package models

import (
	"database/sql"
)

const (
	sqlInsert = `INSERT INTO sessions (AccountID, UUID) VALUES(?, ?)`
	//sqlSingle = `SELECT uuid FROM sessions WHERE id=? AND account_id=?`
	//sqlDelete = `DELETE FROM sessions WHERE id=?`
)

type Session struct {
	Id        int64
	AccountID int64
	UUID      string
}

type SessionModel struct {
	DB *sql.DB
}

func (m SessionModel) SessionInsert(requestSession Session) error {
	_, err := m.DB.Exec(
		sqlInsert,
		requestSession.AccountID,
		requestSession.UUID)
	return err
}
