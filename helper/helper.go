package helper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RespondJSON responds with a JSON payload.
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	byteData, err := json.Marshal(payload)
	if err != nil {
		SugarObj.Error(err.Error())
		return
	}
	err = json.Unmarshal(byteData, &payload)
	if err != nil {
		SugarObj.Error(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// hashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func JSONMarshal(v map[string]string, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckString(statuscreatedat sql.NullString) string {
	statuscreatedat_ := ""
	if statuscreatedat.Valid {
		statuscreatedat_ = statuscreatedat.String
	}
	return statuscreatedat_
}

func TypeCheck(create_user_stat interface{}) string {
	switch create_user_stat.(type) {
	case string:
		return create_user_stat.(string)

	case float64:
		s := fmt.Sprintf("%d", int(create_user_stat.(float64)))
		return s

	default:
		return ""

	}
}

func UUID() string {
	uuid := uuid.New()
	return uuid.String()
}
