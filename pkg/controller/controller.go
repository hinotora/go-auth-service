package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hinotora/go-auth-service/pkg/config"
	"github.com/hinotora/go-auth-service/pkg/logger"
	"github.com/hinotora/go-auth-service/pkg/models"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Data: make(map[string]interface{})}

	response.Status = "ok"
	response.Data["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}

func Login(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	exp := time.Now().Add(time.Duration(config.Instance.Auth.Jwt_time_to_live) * time.Minute)

	key := []byte(config.Instance.Auth.Jwt_secret_key)

	response := models.Response{Data: make(map[string]interface{})}

	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)

	if err != nil {
		w.WriteHeader(500)

		logger.Logger.Println(err)

		response.Status = "error"
	} else {
		response.Status = "ok"
		response.Data["token"] = token
	}

	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func Check(w http.ResponseWriter, r *http.Request) {
	auth_header := strings.Split(r.Header["Authorization"][0], " ")

	response := models.Response{Data: make(map[string]interface{})}

	if len(auth_header) < 2 || auth_header[0] != "Bearer" {
		w.WriteHeader(401)

		response.Status = "error"

		return
	}

	token, err := jwt.Parse(auth_header[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		key := []byte(config.Instance.Auth.Jwt_secret_key)

		return key, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response.Status = "ok"
	} else {
		logger.Logger.Println(err)

		response.Status = "error"

		w.WriteHeader(401)
	}

	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
