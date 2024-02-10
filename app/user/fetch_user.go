package user

import (
	"net/http"
	"secure-sign/config"
	helper "secure-sign/helper"
)

func getUser(w http.ResponseWriter, r *http.Request, username string) (response config.Response) {

	var fetchedUser User

	if helper.IsPatternValid("Email", username) || helper.IsPatternValid("Mobile", username) {

		query := "SELECT id,first_name, last_name, gender, age, salary FROM users WHERE email = $1 OR phone_number = $2"
		row := helper.QueryRow(query, username, username)

		err := row.Scan(&fetchedUser.ID, &fetchedUser.FirstName, &fetchedUser.LastName, &fetchedUser.Gender, &fetchedUser.Age, &fetchedUser.Salary)
		if err != nil {
			response := config.Response{
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
			helper.SugarObj.Error("%+v", response)
			return response
		}

		response = config.Response{
			Status:     "Success",
			StatusCode: http.StatusOK,
			Message:    "User details fetched successfully",
			Data:       fetchedUser,
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
