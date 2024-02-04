package health

import (
	"net/http"
	"secure-sign/config"
	"secure-sign/helper"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	response := config.Response{
		Status:     "Success",
		StatusCode: http.StatusOK,
		Message:    "Healthy",
	}
	helper.SugarObj.Error(response)
	helper.RespondJSON(w, http.StatusBadRequest, response)
	
}
