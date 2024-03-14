package xauth

import (
	"auth/database/permisos"
	"auth/database/usuarios"
	"auth/graph/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaim struct {
	USERID string `json:"id"`
	jwt.StandardClaims
}

type AuthData struct {
	Clains  *JwtCustomClaim
	Usuario *model.Usuario
	TOKEN   string `json:"token"`
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}

func jwtGenerate(userID string, minutos int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		USERID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(minutos)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtValidate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}

type authString string

func AuthMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			if auth == "" || len(auth) <= 7 {
				next.ServeHTTP(w, r)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			validate, err := JwtValidate(auth)
			if err != nil {
				txt := err.Error()
				if !strings.HasPrefix(txt, "token is expired by") {
					next.ServeHTTP(w, r)
					return
				}
			}

			customClaim, _ := validate.Claims.(*JwtCustomClaim)

			data := AuthData{}
			data.Clains = customClaim
			data.TOKEN = auth
			us, er := usuarios.GetById(db, customClaim.USERID)
			if er == nil {
				data.Usuario = us
			}

			ctx := context.WithValue(r.Context(), authString("auth"), &data)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})
	}
}

func GenerateToken(ctx context.Context, userID string) (string, error) {
	tokenduration := os.Getenv("TOKEN_DURATION_MIN")
	duration := 10
	id, err := strconv.Atoi(tokenduration)
	if err == nil {
		duration = id
	}
	return jwtGenerate(userID, duration)
}

func GenerateRefreshToken(ctx context.Context, userID string) (string, error) {
	tokenduration := os.Getenv("TOKENREFRESH_DURATION_MIN")
	duration := 60
	id, err := strconv.Atoi(tokenduration)
	if err == nil {
		duration = id
	}
	return jwtGenerate(userID, duration)
}

func CtxValue(ctx context.Context, db *sql.DB, metodo string) (*AuthData, error) {
	str := authString("auth")
	algo := ctx.Value(str)
	if algo == nil {
		return nil, errors.New("proporcione un token")
	}
	clains, _ := algo.(*AuthData)
	if clains == nil {
		return nil, errors.New("debes iniciar session")
	}
	validate, err := JwtValidate(clains.TOKEN)
	if err != nil || !validate.Valid {
		txt := err.Error()
		if strings.HasPrefix(txt, "token is expired by") {
			txt = strings.Replace(txt, "token is expired by", "Su sessión expiró hace ", 1)
			return nil, errors.New(txt)
		} else {
			return nil, errors.New(txt)
		}
	}
	if !clains.Usuario.Estado {
		return nil, errors.New("tu cuenta se encuentra suspendida")
	}

	if len(metodo) > 0 {
		err = permisos.VerificarPermiso(db, clains.Usuario.ID, metodo)
		if err != nil {
			return nil, err
		}
	}

	return clains, nil
}
