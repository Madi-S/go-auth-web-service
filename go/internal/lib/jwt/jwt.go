package jwt

import (
	"go-auth-service/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":    user.ID,
		"email":  user.Email,
		"app_id": app.ID,
		"exp":    time.Now().Add(duration).Unix(),
	})

	return token.SignedString([]byte(app.Secret))
}
