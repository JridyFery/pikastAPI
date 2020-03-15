package helpers

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

// IsAuthorized func
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header.Get("Authorization")[7:] != "" {

            token, err := jwt.Parse(r.Header.Get("Authorization")[7:], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                role := token.Claims.(jwt.MapClaims)["role"]
                if roleAuthorization(role.(string),r.URL.Path) {
                    endpoint(w, r)
                }else {
                    fmt.Fprintf(w, "Not Authorized")
                }
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

func roleAuthorization(role string, endpoint string ) bool {
    var roles Roles
    file, _ := ioutil.ReadFile("access.json")
	json.Unmarshal([]byte(file), &roles)
	if role == "user" {
        for _, r := range roles.User {
            if r == endpoint {
                return true
            }
        }
    }
    if role == "admin" {
        for _, r := range roles.Admin {
            if r == endpoint {
                return true
            }
        }
    }
    if role == "super admin" {
        for _, r := range roles.SuperAdmin {
            if r == endpoint {
                return true
            }
        }
    }
    return false
}

// Roles struct
type Roles struct {
    Admin       []string    `json:"admin"`
    SuperAdmin  []string    `json:"superAdmin"`
    User        []string    `json:"user"`
}