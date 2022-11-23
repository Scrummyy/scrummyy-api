package models

import (
	"errors"

	db "github.com/Scrummyy/scrummyy-api/configs"
	datatype "github.com/Scrummyy/scrummyy-api/internal/datatypes"

	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID       string `db:"id, primarykey, autoincrement" json:"id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
	Name     string `db:"name" json:"name"`
	// Projects  []sql.NullString `db:"projects" json:"projects"`
	// Workspace string           `db:"workspace" json:"workspace"`
	UpdatedAt string `db:"updated_at" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
}

// UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

// Login ...
func (m UserModel) Login(form datatype.LoginForm) (user User, token Token, err error) {

	err = db.GetDB().SelectOne(&user, "SELECT id, name, email, username, created_at FROM auth.users WHERE email=LOWER($1) LIMIT 1", form.Email)

	if err != nil {
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

// Register ...
func (m UserModel) Register(form datatype.RegisterForm) (user User, err error) {
	getDb := db.GetDB()

	//Check if the user exists in database
	checkUser, err := getDb.SelectInt("SELECT count(id) FROM auth.users WHERE email=LOWER($1) LIMIT 1", form.Email)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	if checkUser > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	//Create the user and return back the user ID
	err = getDb.QueryRow("INSERT INTO auth.users(email, password, name, username) VALUES($1, $2, $3, $4) RETURNING id", form.Email, string(hashedPassword), form.Name, form.Username).Scan(&user.ID)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	user.Name = form.Name
	user.Email = form.Email

	return user, err
}

// One ...
func (m UserModel) One(userID int64) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM auth.users WHERE id=$1 LIMIT 1", userID)
	return user, err
}
