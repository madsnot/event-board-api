package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/madsnot/event-board-api/internal/domain/models"
	jwt "github.com/madsnot/event-board-api/pkg/token"
)

const (
	bearerMethod    = "Bearer"
	headerAuthorize = "Authorization"
)

func AuthMiddleware(next http.HandlerFunc, tokenizer *jwt.Tokenizer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := checkAuthorization(r.Context(), bearerMethod)
		if err != nil {
			writeError(w, err)
		}

		claims := models.Claims{}

		if err = tokenizer.UnmarshalJWT(token, &claims); err != nil {
			writeError(w, err)
		}

		if claims.IsValidAt(time.Now().UTC()) {
			writeError(w, ErrJWTExpired)
		}

		next(w, r)
	}
}

func checkAuthorization(ctx context.Context, scheme string) (string, error) {
	value := ctx.Value(headerAuthorize).(string)

	if value == "" {
		return "", fmt.Errorf("request unauthenticated with %s", scheme)
	}

	splits := strings.SplitN(value, " ", 2)
	if len(splits) < 2 {
		return "", errors.New("bad authorization string")
	}

	if !strings.EqualFold(splits[0], scheme) {
		return "", fmt.Errorf("request unauthenticated with %s", scheme)
	}

	return splits[1], nil
}
