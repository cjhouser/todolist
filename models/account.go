package models

import (
	"database/sql"
)

const (
	sqlAccountInsert = `INSERT INTO accounts (username, password, salt) VALUES(?, ?)`
	sqlAccountSingle = `SELECT id, username, password, salt FROM accounts WHERE id=?`
	sqlAccountUpdate = `UPDATE accounts SET password=?, salt=? WHERE id=?`
	sqlAccountDelete = `DELETE FROM accounts WHERE id=?`
)

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
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

func (m AccountModel) AccountSingle(requestAccount Account) (responseAccount *Account, err error) {
	row := m.DB.QueryRow(
		sqlAccountSingle,
		requestAccount.Id)
	responseAccount = &Account{}
	err = row.Scan(&responseAccount.Id, &responseAccount.Username, &responseAccount.Password, &responseAccount.Salt)
	if err != nil {
		return nil, err
	}
	return responseAccount, nil
}

func (m AccountModel) AccountUpdate(requestAccount Account) error {
	_, err := m.DB.Exec(
		sqlAccountUpdate,
		requestAccount.Password,
		requestAccount.Salt,
		requestAccount.Id)
	return err
}

func (m AccountModel) AccountDelete(requestAccount Account) error {
	_, err := m.DB.Exec(
		sqlAccountDelete,
		requestAccount.Id)
	return err
}
