package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"subway/server/response"

	"github.com/gin-gonic/gin"
)

func RequestDecoding(context *gin.Context, data interface{}) {

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
}

func SetHeader(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

}

func GetTokenFromAuthHeader(context *gin.Context) (string, error) {
	token := strings.Split(context.Request.Header["Authorization"][0], " ")[1]
	if token == "" {
		return "", errors.New("token not found")
	}
	return token, nil
}
