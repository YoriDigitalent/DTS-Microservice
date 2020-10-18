package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"

	"github.com/YoriDigitalent/DTS-Microservice/auth-service/database"
	"github.com/YoriDigitalent/DTS-Microservice/auth-service/util"
	"github.com/wskurniawan/intro-microservice/utils"
)

type AuthDB struct {
	Db *gorm.DB
}

func ValidateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")

	res, err := database.ValidateAuth(authToken, db.Db); if err != nil {
		utils.WrapAPIError( w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPISuccess(w, r, "Auth Success", http.StatusOK)
}

func (db *AuthDB) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "Can't read body", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w, r, "error Unmarshal: "+err.Error(), http.StatusInternalServerError)
		return
	}

	signup.Token = util.IdGenerator()

	err = signup.SignUp(db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
	return

	//TO DO buat SignUp
}

func (db *AuthDB) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "Can't read body", http.StatusBadRequest)
		return
	}

	var login database.Auth

	err = json.Unmarshal(body, &login)
	if err != nil {
		utils.WrapAPIError(w, r, "error Unmarshal: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := login.Login(db, Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPIData(w, r, database.Auth{
		Username: res.Username,
		Token:    res.Token,
	}, http.StatusOK, "Success")

	//TO DO buat Login
}
