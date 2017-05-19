package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipe/models"

	"strconv"

	"recipe/helpers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type usersResponse struct {
	Status int           `json:"status"`
	Data   []models.User `json:"data"`
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		UserName:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	decoder := json.NewDecoder(r.Body)

	decoderErr := decoder.Decode(&user)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	if user.FirstName == "" || user.LastName == "" ||
		(user.UserName == "" || len(user.UserName) < 3) || !helpers.IsValidEmail(user.Email) ||
		(user.Password == "" || len(user.Password) < 6) {
		errMsg := models.User{
			FirstName: "first name is required",
			LastName:  "Last  name is required",
			UserName:  "Username is required and should be more that 3 characters",
			Email:     "Email is required and should be valid",
			Password:  "Password is required and should be more than 6 characters",
		}

		type Error struct {
			Status  int
			Message interface{}
		}
		newError := Error{Status: http.StatusBadRequest, Message: errMsg}
		response, _ := json.Marshal(newError)
		helpers.ResponseWriter(w, http.StatusBadRequest, string(response))
		return
	}
	_, dbErr := models.CreateUser(&user)

	if dbErr != nil {
		helpers.ServerError(w, dbErr)
		return
	}
	helpers.StatusOk(w, user)
}

// GetUser Gets all users and sends the data as response
// to the requesting user
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	user, err := models.GetUser(id)
	if err != nil {
		helpers.StatusNotFound(w, err)
		return
	}
	helpers.StatusOk(w, user)
}

// GetAllUsers Gets all users and sends the data as response
// to the requesting user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUser()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	response := usersResponse{
		Status: http.StatusOK,
		Data:   users,
	}
	result, _ := json.Marshal(response)

	helpers.ResponseWriter(w, http.StatusOK, string(result))
}

//UpdateUser updates user's detail
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		UserName:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	decoder := json.NewDecoder(r.Body)

	decoderErr := decoder.Decode(&user)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := models.UpdateUser(id, &user)
	fmt.Println(err.Error())
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}
	helpers.StatusOk(w, user)
}

//DeleteUser deletes a user detail
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	_, err := models.DeleteUser(id)
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}
	helpers.ResponseWriter(w, http.StatusOK, "User deleted")
}