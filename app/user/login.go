package user

import (
	"database/sql"
	"net/http"
	"secure-sign/config"
	helper "secure-sign/helper"

	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request, credentials LoginRequest) (response config.Response) {
	response = config.Response{}

	// Sanitize input
	credentials.Username = helper.SanitizeInput(credentials.Username)
	credentials.Password = helper.SanitizeInput(credentials.Password)

	// Validate if username is either email or phone number
	if helper.IsPatternValid("Email", credentials.Username) || helper.IsPatternValid("Mobile", credentials.Username) {

		query := `
			SELECT password FROM users
			WHERE email = $1 OR phone_number = $1
		`
		// Execute the query using a prepared statement
		row := helper.QueryRow(query, credentials.Username)
		var fetchedPassword string
		err := row.Scan(&fetchedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				// User with the provided email or phone number not found
				response = config.Response{
					Status:     "success",
					StatusCode: http.StatusOK,
					Message:    "Invalid Credentials",
				}
				return response
			}

			response = config.Response{
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid username",
			}
			helper.SugarObj.Error(err)
			return response
		}

		// Compare the fetched password with the provided password
		err = bcrypt.CompareHashAndPassword([]byte(fetchedPassword), []byte(credentials.Password))
		if err != nil {
			response = config.Response{
				Status:     "success",
				StatusCode: http.StatusOK,
				Message:    "Invalid credentials",
			}
			helper.SugarObj.Error(err)
			return response
		}

		// Respond with a success message
		response = config.Response{
			Status:     "success",
			StatusCode: http.StatusOK,
			Message:    "Login successful",
		}
		helper.SugarObj.Info("%+v", response)
		return response
	}

	response = config.Response{
		Status:     "Bad Request",
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid username",
	}
	helper.SugarObj.Error("%+v", response)
	return response
}
