package user

import (
	"net/http"
	"secure-sign/config"
	helper "secure-sign/helper"

	"github.com/gorilla/mux"
)

// User represents the user model.
type User struct {
	ID          string `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Age         int    `json:"age,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Salary      int    `json:"salary,omitempty"`
	Password    string `json:"password,omitempty"`
}

// LoginRequest represents the login credentials.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterHandler handles the registration endpoint.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	var response config.Response
	// Decode the request body into the newUser struct
	err := helper.DecodeRequest(r, &newUser)
	if err != nil {
		response = config.Response{
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		helper.SugarObj.Error(err)
		helper.RespondJSON(w, http.StatusBadRequest, response)
		return
	}

	response = registration(w, r, newUser)
	helper.RespondJSON(w, response.StatusCode, response)

}

// LoginHandler handles the login endpoint.

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credentials LoginRequest

	response := config.Response{}
	// Decode the request body into the credentials struct

	err := helper.DecodeRequest(r, &credentials)
	if err != nil {
		response := config.Response{
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		helper.SugarObj.Error(err)
		helper.RespondJSON(w, http.StatusBadRequest, response)
		return
	}

	response = login(r.Context(), credentials)
	helper.RespondJSON(w, response.StatusCode, response)

}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var credentials LoginRequest
// 	// Get database connection from context
// 	db := r.Context().Value("db").(*sql.DB)

// 	fmt.Printf("credentials: %+v\n", credentials)
// 	fmt.Printf("db: %+v\n", db)

// 	// Decode the request body into the credentials struct
// 	err := helper.DecodeRequest(r, &credentials)
// 	if err != nil {
// 		response := config.Response{
// 			Status:     "Bad Request",
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 		}
// 		fmt.Println("failed:: ", err)
// 		helper.SugarObj.Error(err)
// 		helper.RespondJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	response := login(r.Context(), db, credentials)
// 	helper.RespondJSON(w, response.StatusCode, response)
// }

// GetUserHandler handles the fetch user endpoint.
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	response := getUser(w, r, username)
	// Respond with the fetched user details

	helper.RespondJSON(w, http.StatusOK, response)
}
