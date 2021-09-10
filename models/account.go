package models

import (
	"database/sql"
)

const (
	sqlAccountInsert = `INSERT INTO accounts (username, password, salt) VALUES(?, ?, ?)`
	sqlAccountSingle = `SELECT (username, password, salt) FROM accounts WHERE id=?`
)

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
	Salt     []byte `json:"salt"`
}

type AccountModel struct {
	DB *sql.DB
}

func (m AccountModel) AccountInsert(requestAccount Account) error {
	_, err := m.DB.Exec(
		sqlAccountInsert,
		requestAccount.Username,
		requestAccount.Password,
		requestAccount.Salt)
	return err
}

func (m AccountModel) AccountSingle(requestAccount Account) error {
	return nil
}
