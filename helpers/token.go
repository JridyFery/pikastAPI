package helpers

import(
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"os"
)

var mySigningKey = []byte(os.Getenv("MY_JWT_TOKEN"))

// GenerateJWT func
func GenerateJWT(userName string, role string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = userName
    claims["role"] = role
    claims["exp"] = time.Now().Add(time.Hour * 999999).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}