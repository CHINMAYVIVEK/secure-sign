package user

import (
	"fmt"
	"net/http"
	"secure-sign/config"
	helper "secure-sign/helper"
)

func registration(w http.ResponseWriter, r *http.Request, newUser User) (response config.Response) {

	// Validate the required fields
	if err := validateRequiredFields(newUser); err != nil {
		response = config.Response{
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		helper.SugarObj.Error(err)
		return response
	}

	// Check if email or phone number already exists
	query := `
		SELECT
			COUNT(*) FILTER (WHERE email = $1) as email_count,
			COUNT(*) FILTER (WHERE phone_number = $2) as phone_count
		FROM users
	`
	var emailCount, phoneCount int

	err := helper.QueryRow(query, newUser.Email, newUser.PhoneNumber).Scan(&emailCount, &phoneCount)
	if err != nil {
		response = config.Response{
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message:    "Error checking email and phone number existence",
		}
		helper.SugarObj.Error(err)
		return response
	}

	// Check if email already exists
	if emailCount > 0 {
		response = config.Response{
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message:    "Email already exists",
		}
		helper.SugarObj.Error("%+v", response)
		return response
	}

	// Check if phone number already exists
	if phoneCount > 0 {
		response = config.Response{
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message:    "Phone number already exists",
		}
		helper.SugarObj.Error("%+v", response)
		return response
	}

	// Hash the password before storing it
	hashedPassword, err := helper.HashPassword(newUser.Password)
	if err != nil {
		response = config.Response{
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Message:    "Error hashing password",
		}
		helper.SugarObj.Error("%+v", response)
		return response

	}
	newUser.Password = hashedPassword
	newUser.ID = helper.UUID()
	// Query to insert a new user record
	query = `
		INSERT INTO users (id,first_name, last_name, gender, age, email, phone_number, salary, password)
		VALUES ($1,$2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	_, err = helper.Exec(query, newUser.ID, newUser.FirstName, newUser.LastName, newUser.Gender, newUser.Age, newUser.Email, newUser.PhoneNumber, newUser.Salary, newUser.Password)

	if err != nil {
		response = config.Response{
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		helper.SugarObj.Error("%+v", response)
		return response
	}

	response = config.Response{
		Status:     "Success",
		StatusCode: http.StatusCreated,
		Message:    "User Created Successfully",
	}
	helper.SugarObj.Error("%+v", response)
	return response

}

func validateRequiredFields(user User) error {

	user.FirstName = helper.SanitizeInput(user.FirstName)
	user.LastName = helper.SanitizeInput(user.LastName)

	user.Email = helper.SanitizeInput(user.Email)
	user.Password = helper.SanitizeInput(user.Password)

	if user.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if user.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	if !helper.IsPatternValid("Email", user.Email) {
		return fmt.Errorf("email is not valid")
	}

	if !helper.IsPatternValid("Mobile", user.PhoneNumber) {
		return fmt.Errorf("phone number is not valid")
	}
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	if user.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
