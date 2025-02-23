package middleware

import (
	"Find-Backend/core/common"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() fiber.Handler {
	excludedPaths  := map[string]bool  {
		"/auth/login": true,
		"/auth/register": true,
	}

	return func (ctx *fiber.Ctx) error {
		if _, excluded := excludedPaths[ctx.Path()]; excluded {
			return ctx.Next()
		}
	
		authHeader := ctx.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return ctx.Status(fiber.StatusUnauthorized).JSON(common.NewFailedResponse("Tidak memiliki Hak Akses", "Silahkan login terlebih dahulu"))
		}
	
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	
		key := os.Getenv("SECRET_KEY")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("metode autentikasi tidak diketahui")
			}
	
			return []byte(key),  nil
		})
		if err != nil || !token.Valid {
			switch err.Error() {
				case "metode autentikasi tidak diketahui":
					
				case "Token is expired":
					err = errors.New("token telah kadaluarsa")
	
				default:
					log.Printf("Error Autentikasi: %v", err.Error())
					err = errors.New("token tidak valid")
			}
	
			return ctx.Status(fiber.StatusUnauthorized).JSON(common.NewFailedResponse("Autentikasi gagal", err.Error()))
		}
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id, exist := claims["id"]
			if !exist || id == "" {
				err = errors.New("format token tidak valid")
			}

			rawId, ok := id.(string)
			if !ok {
				log.Printf("Error saat mengambil id, tipe data id : %T, nilai id : %v", id, id)
				err = errors.New("format token tidak valid")
			}

			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(common.NewFailedResponse("Autentikasi gagal", err.Error()))
			}
	
			ctx.Locals("id", rawId)
		} else {
			return ctx.Status(fiber.StatusUnauthorized).JSON(common.NewFailedResponse("Autentikasi gagal", "Format token tidak valid"))
		}
	
		return ctx.Next()
	}

}