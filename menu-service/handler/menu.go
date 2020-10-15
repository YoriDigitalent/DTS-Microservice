package handler

import (
	"github.com/YoriDigitalent/DTS-Microservice/utils"
	"net/http"
)

func AddMenu(w http.ResponseWriter, r *http.Request)) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
