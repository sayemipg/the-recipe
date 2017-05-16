package models

import (
	"database/sql"
	"fmt"
	"log"
	"recipe/helpers"
)

// User data to be sent
// When request is made to the server
type User struct {
	ID         string `json:"id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	UserName   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	ProfilePic string `json:"profile_pic,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

// GetAll gets all user
func GetAllUser() []User {
	db := DB()
	users := []User{}

	rows, err := db.Query(`SELECT id, first_name, last_name, 
		username, email, profile_pic, created_at, updated_at FROM users`)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		user := User{}

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName,
			&user.Email, &user.ProfilePic, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err.Error())
		}
		users = append(users, user)
		fmt.Println(user)
	}
	return users
}

// Get gets a single user
func GetUser(id int) (User, error) {
	db := DB()
	user := User{}
	db.Close()

	err := db.QueryRow("SELECT first_name, last_name, username, email, profile_pic FROM users where id = ? ", id).Scan(&user.FirstName, &user.Email, &user.LastName, &user.UserName, &user.ProfilePic)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows with that id found", err.Error())
		errMsg := fmt.Errorf("user with (id %d) does not exist", id)
		return user, errMsg
	case err != nil:
		log.Fatal("An error occurred", err.Error())
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return user, errMsg
	default:
		return user, nil
	}
}

// CreateUser creates a new user
func (u User) CreateUser(db *sql.DB) {
	sql := `INSERT INTO users(first_name, last_name, email, username,
				profile_pic ) VALUES(?, ?, ?, ?, ?)`
	row, err := db.Exec(sql, u.FirstName, u.LastName, u.Email, u.UserName, u.ProfilePic)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row)
}

// DeleteUser user from database
// NB: this method wipes all user details
// recipe ans ingredients inclusive
func DeleteUser(id int) (bool, error) {
	db := DB()

	sql := `DELETE * FROM users WHERE id = ?`
	row, err := db.Exec(sql, id)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(row)

	return true, nil
}

// UpdateUser user details base on the values sent
// takes the user id and user struct containing
// details to be update
func UpdateUser(id int, user *User) (bool, error) {
	db := DB()

	userValues := map[string]string{}

	if user.FirstName != "" {
		userValues["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		userValues["last_name"] = user.LastName
	}
	if user.UserName != "" {
		userValues["username"] = user.UserName
	}
	if user.ProfilePic != "" {
		userValues["profile_pic"] = user.ProfilePic
	}
	if user.Email != "" {
		userValues["email"] = user.Email
	}

	query := helpers.UpdateBuilder(userValues)

	println(query)
	rows, err := db.Exec(query, id)
	if err != nil {
		panic(err.Error())
	} else {
		println(rows)
	}

	return true, nil
}