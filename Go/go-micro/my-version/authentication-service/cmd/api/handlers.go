package main

import (
	"authentication/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)
func (app *Config) InsertUser(w http.ResponseWriter, r *http.Request) {
	theUser := data.Models {
		User: data.User{
			ID:        1,
			Email:     "mhu@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "password",
			Active:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Insert the user into the database
	_, err := theUser.User.Insert()
	if err != nil {
		// Handle the error, for example, by returning an error response
		app.writeJSON(w, http.StatusInternalServerError, "Error inserting user")
		return
	}

	// log authentication
	err = app.logRequest("authentication", fmt.Sprintf("%s logged in", theUser.User.Email))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// Respond with a success message
	app.writeJSON(w, http.StatusOK, "User inserted successfully")
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email string `json:"email"`
		Password string `json:"password"`	
	}
	err := app.readJSON(w, r, &requestPayload)
	log.Println("email is " + requestPayload.Email)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	log.Printf("user Password is %s ", user.Password)
	log.Printf("user err is %s ", err)

	log.Printf("request password is %s ", requestPayload.Password)

	// valid := user.Password == requestPayload.Password
	valid, err := user.PasswordMatches(requestPayload.Password)    // plain password compared to hash encoded password
	if !valid {
		log.Println("user valid is %s ", valid)
		// log.Println("user err is %s ", err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse {
		Error: false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data: user,
	}
	log.Printf("Response from auth service %s\n", payload)

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"
	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}