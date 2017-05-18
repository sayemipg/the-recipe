package helpers

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// Multiplier multiplies all
// integer parameters passed
// into this function
func Multiplier(args ...int) int {
	ans := 0
	for _, value := range args {
		ans += value
	}
	return ans
}

// PrintErr handles unexpected
// errors that occurs
func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// RespondMessage converts error message to a
// valid string to be sent as response
func RespondMessage(err string, status int) string {
	type Response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	newErr := Response{
		Status:  status,
		Message: err,
	}

	response, _ := json.Marshal(newErr)
	return string(response)
}

// RespondMessages converts error messages to a
// valid string to be sent as response
func RespondMessages(err interface{}, status int) string {
	type Response struct {
		Status  int         `json:"status"`
		Message interface{} `json:"message"`
	}
	newErr := Response{
		Status:  status,
		Message: err,
	}

	response, err := json.Marshal(newErr)
	return string(response)
}

// RespondWithData converts error message to a
// valid string to be sent as response
func RespondWithData(message string, status int, data interface{}) string {
	type Response struct {
		Status  int         `json:"status"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data"`
	}
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	result, _ := json.Marshal(response)
	return string(result)
}

//IsValidEmail is
func IsValidEmail(email string) bool {
	if m, _ := regexp.MatchString(`([\w\.\_]{2,})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}
