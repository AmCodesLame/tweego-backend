package middleware

import (
	"fmt"
	"os"
	"strings"

	beecontext "github.com/beego/beego/v2/server/web/context"
	jwt "github.com/golang-jwt/jwt/v5"
)

func Authenticate(ctx *beecontext.Context) {
	// cookieToken := ctx.GetCookie("bearer")
	authorizationHeader := ctx.Input.Header("Authorization")
	if authorizationHeader == "" {
		ctx.Output.SetStatus(401)
		_ = ctx.Output.Body([]byte("Unauthorized\n"))
		return
	}
	token1 := strings.Split(authorizationHeader, " ")[1]
	fmt.Println(token1)
	// cookieToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZWRBdCI6MTcwMjQ3ODkyNSwidGVzdCI6InRoaXMgaXMgYSB0ZXN0In0.hWOIrg8hnKJ7-244nXOG5Rg0lFwDpf8ygFCduWx50Mc"
	token, err := jwt.Parse(token1, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte([]byte(os.Getenv("HMACKEY"))), nil
	})
	if err != nil {
		ctx.Output.SetStatus(401)
		fmt.Println("err is ", err)
		_ = ctx.Output.Body([]byte("Unauthorized\n"))
		return
	}
	if !token.Valid {
		ctx.Output.SetStatus(401)
		_ = ctx.Output.Body([]byte("Unauthorized\n"))
		return
	}

	// if claims, ok := token.Claims.(jwt.MapClaims); ok {
	// 	fmt.Println(claims["foo"], claims["nbf"])
	// } else {
	// 	fmt.Println(err)
	// }
}
