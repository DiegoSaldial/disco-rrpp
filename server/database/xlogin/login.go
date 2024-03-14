package xlogin

import (
	"auth/database/usuarios"
	"auth/database/xauth"
	"auth/graph/model"
	"context"
	"database/sql"
	"errors"
	"strings"
)

func Login(db *sql.DB, data model.NewLogin) (*model.ResponseLogin, error) {
	us, err := usuarios.GetByUserPass(db, data.Username, data.Password)
	if err != nil {
		return nil, err
	}
	if !us.Estado {
		return nil, errors.New("usuario no activo")
	}

	ctx := context.Background()
	token, err := xauth.GenerateToken(ctx, us.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := xauth.GenerateRefreshToken(ctx, us.ID)
	if err != nil {
		return nil, err
	}

	res := model.ResponseLogin{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return &res, nil
}

func RefreshToken(token, refreshToken string) (string, error) {
	userid, err := validateTokens(token, false)
	if err != nil {
		return "", err
	}

	_, err = validateTokens(refreshToken, true)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	tok, err := xauth.GenerateToken(ctx, userid)
	if err != nil {
		return "", err
	}
	return tok, nil
}

func validateTokens(t string, isrefresh bool) (string, error) {
	validate, err := xauth.JwtValidate(t)
	customClaim, _ := validate.Claims.(*xauth.JwtCustomClaim)

	if err != nil || !validate.Valid {
		txt := err.Error()
		expire := strings.HasPrefix(txt, "token is expired by")
		if expire && isrefresh {
			txt = strings.Replace(txt, "token is expired by", "Su sessión expiró hace ", 1)
			return "", errors.New(txt)
		} else if expire && !isrefresh {
			return customClaim.USERID, nil
		} else {
			return "", errors.New(txt)
		}
	}

	return customClaim.USERID, nil
}
