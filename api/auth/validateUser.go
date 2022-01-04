package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/matthieurobert/amp/api/entity"
	"github.com/shaj13/go-guardian/v2/auth"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	user, err := entity.GetUserByUsername(userName)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return auth.NewDefaultUser(user.Username, string(rune(user.Id)), nil, nil), nil
}
