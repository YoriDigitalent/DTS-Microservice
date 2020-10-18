package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/YoriDigitalent/DTS-Microservice/menu-service/config"
	"github.com/YoriDigitalent/DTS-Microservice/menu-service/entity"
	"github.com/YoriDigitalent/DTS-Microservice/utils"
	"github.com/gorilla/context"
)

type AuthHandler struct {
	config config.Auth
}

//menjalankan validasi terlebih dahulu -> baru nextHandler
func (handler *AuthHandler) ValidateAdmin(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest(http.MethodPost, handler.config.Host+"/validate-admin", nil)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header

		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody, err := ioutil.ReadAll(authResponse.Body)

		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var authResult entity.AuthResponse
		err = json.Unmarshal(responseBody, &authResult)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		if authResponse.StatusCode != http.StatusOK {
			utils.WrapAPIError(w, r, authResult.ErrorDetails, authResponse.StatusCode)
			return

		}

		context.Set(r, "user", authResult.Data.Username)

		nextHandler(w, r)

	}

}
