package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Miguelburitica/goHttpServer/src/controllers"
	jwt "github.com/dgrijalva/jwt-go"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)
		fmt.Println(r.URL)
		fmt.Println(r.Header)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "userID",
			"iat": time.Now().Unix(),
		})

		tokenString, err := token.SignedString([]byte("secret"))

		if err != nil {
			panic(err)
		}
		
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/users", Auth(controllers.ShowHome))

	http.ListenAndServe(":3000", nil)
}