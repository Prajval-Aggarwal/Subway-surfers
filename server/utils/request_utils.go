package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequestDecoding decoed the request body and stores it in data interface
func RequestDecoding(context *gin.Context, data interface{}) error {

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		return err
	}
	return nil
}

// SetHeader sets the header to value of application/json
func SetHeader(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

}

// GetTokenFromAuthHeader returns the token from auth header
func GetTokenFromAuthHeader(context *gin.Context) (string, error) {
	token := strings.Split(context.Request.Header["Authorization"][0], " ")[1]
	if token == "" {
		return "", errors.New("token not found")
	}
	return token, nil
}
