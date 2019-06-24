// Copyright 2019 pandora Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ielepro/pandora/module/user"
	"github.com/ielepro/pandora/render"
	reqApi "github.com/ielepro/pandora/router/route/api"
	"github.com/ielepro/pandora/util/goslice"
	"strings"
)

func ApiPriv() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("_syd_identity")

		path := strings.TrimSpace(c.Request.URL.Path)
		if len(path) < 4 {
			respondWithError(c, render.CODE_ERR_NO_PRIV, "request path is too short")
			return
		}
		path = path[4:]

		if path == reqApi.LOGIN {
			c.Next()
			return
		}

		if token == "" {
			respondWithError(c, render.CODE_ERR_NO_LOGIN, "no login")
			return
		}

		login := &user.Login{
			Token: token,
		}
		if err := login.ValidateToken(); err != nil {
			respondWithError(c, render.CODE_ERR_NO_LOGIN, err.Error())
			return
		}

		//priv check
		role := &user.Role{
			ID: login.RoleId,
		}
		if err := role.Detail(); err != nil {
			respondWithError(c, render.CODE_ERR_APP, err.Error())
			return
		}

		c.Set("user_id", login.UserId)
		c.Set("username", login.Username)
		c.Set("email", login.Email)
		c.Set("truename", login.Truename)
		c.Set("mobile", login.Mobile)
		c.Set("role_name", role.Name)
		c.Set("privilege", role.Privilege)

		needLoginReq := []string{
			reqApi.LOGIN_STATUS,
			reqApi.LOGOUT,
			reqApi.MY_USER_SETTING,
			reqApi.MY_USER_PASSWORD,
		}
		if goslice.InSliceString(path, needLoginReq) {
			c.Next()
			return
		}

		if pass := user.CheckHavePriv(path, role.Privilege); !pass {
			respondWithError(c, render.CODE_ERR_NO_PRIV, "no priv")
			return
		}

		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message string) {
	render.CustomerError(c, code, message)
	c.Abort()
}
