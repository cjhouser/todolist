package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CJHouser/tasklist/models"
)

func (env *accountEnvironment) signUp(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account models.Account
	json.Unmarshal(reqBody, &account)
	err := env.accounts.AccountInsert(account)
	if err != nil {
		log.Printf("createTask: %v\n", err)
	}
}
