package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CJHouser/tasklist/models"
	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
)

func (env *accountEnvironment) signUp(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account models.Account
	json.Unmarshal(reqBody, &account)
	account.Salt = []byte(uniuri.NewLen(8))
	hash, err := bcrypt.GenerateFromPassword(append(account.Password, account.Salt...), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("createAccount: %v\n", err)
		return
	}
	account.Password = hash
	err = env.AccountInsert(account)
	if err != nil {
		log.Printf("createAccount: %v\n", err)
	}
}

func (env *accountEnvironment) signIn(w http.ResponseWriter, r *http.Request) {
	log.Println("not implemented")
}
