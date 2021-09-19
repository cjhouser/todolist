package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CJHouser/tasklist/models"
	"github.com/dchest/uniuri"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (env *accountEnvironment) signUp(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account models.Account
	json.Unmarshal(reqBody, &account)
	account.Salt = []byte(uniuri.NewLen(8))
	hash, err := bcrypt.GenerateFromPassword(append(account.Password, account.Salt...), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("signUp: %v\n", err)
		return
	}
	account.Password = hash
	err = env.AccountInsert(account)
	if err != nil {
		log.Printf("signUp: %v\n", err)
	}
}

func (env *accountEnvironment) signIn(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestAccount models.Account
	var dbAccount models.Account
	json.Unmarshal(reqBody, &requestAccount)
	accountID, err := env.AccountSingle(dbAccount)
	if err != nil {
		log.Printf("signIn: %v\n", err)
	}
	// need to check if something was actually returned from db
	err = bcrypt.CompareHashAndPassword(dbAccount.Password, append(requestAccount.Password, requestAccount.Salt...))
	if err != nil {
		log.Printf("signIn: %v\n", err)
	}
	session := models.Session{
		AccountID: accountID,
		UUID:      uuid.NewString(),
	}
	err = env.SessionInsert(session)
	if err != nil {
		log.Printf("signIn: %v\n", err)
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  session.UUID,
		MaxAge: 300,
	}
	http.SetCookie(w, cookie)
}
