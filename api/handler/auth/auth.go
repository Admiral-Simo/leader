package auth

import (
	"errors"
	"fmt"
	"net/http"
	"server/api/cookies"
	"server/api/handler"
	"server/api/web/pages/authtempl"
	"server/store"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Routes(h handler.Handler) {
	h.App.POST("/signup", func(ctx *gin.Context) {
		// Sleep for preventing malicious users
		time.Sleep(time.Millisecond * 200)
		// Get data
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		// Check if user already exists
		_, err := h.Store.GetUserByEmail(ctx, email)
		if err == nil {
			fmt.Println("user already exists.")
			return
		}
		// Check for credentials
		errorForm := checkForCredentials(name, email, password)
		if len(errorForm) > 0 {
			authtempl.SignUp(errorForm).Render(ctx, ctx.Writer)
			return
		}
		// Create User
		user, err := h.Store.CreateUser(ctx, store.CreateUserParams{
			Name:     name,
			Email:    email,
			Password: password,
		})
		// Checking for creation errors
		if err != nil {
			fmt.Println(err)
			return
		}
		// this is the safe part
		tokenString, err := cookies.GenerateJWT(user.ID)
		if err != nil {
			panic("unable to generate jwt")
		}
		ctx.SetCookie("auth_cookie", tokenString, 3600*24*3, "/", "", false, true)
		// TODO: return a cookie that has a json web token that has the userID
		ctx.Redirect(http.StatusSeeOther, "/")
	})
	h.App.POST("/login", func(ctx *gin.Context) {
		// Sleep for preventing malicious users
		time.Sleep(time.Millisecond * 200)
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
		user, err := h.Store.GetUserByEmail(ctx, email)
		if err != nil {
			errorMessage := errors.New("email or password is incorrect.")
			authtempl.Login(errorMessage).Render(ctx, ctx.Writer)
			return
		}
		if user.Password != password {
			errorMessage := errors.New("email or password is incorrect.")
			authtempl.Login(errorMessage).Render(ctx, ctx.Writer)
			return
		}
		// this is the safe part
		tokenString, err := cookies.GenerateJWT(user.ID)
		if err != nil {
			panic("unable to generate jwt")
		}
		ctx.SetCookie("auth_cookie", tokenString, 3600*24*3, "/", "", false, true)
		// TODO: return a cookie that has a json web token that has the userID
		ctx.Redirect(http.StatusSeeOther, "/")
	})
	h.App.GET("/logout", func(ctx *gin.Context) {
		ctx.SetCookie("auth_cookie", "", -1, "/", "", false, true)
		ctx.Redirect(http.StatusSeeOther, "/about")
	})
}

func checkForCredentials(name string, email string, pass string) map[string]string {
	formErrors := make(map[string]string)
	if len(strings.Fields(name)) < 2 {
		formErrors["name"] = "please enter your first_name and last_name."
	}
	if len(email) < 3 || !strings.Contains(email, "@") || !strings.Contains(email[strings.Index(email, "@"):], ".") {
		formErrors["name"] = "email needs to be valid (e.g. example@gmail.com)."
	}
	if len(pass) < 8 {
		formErrors["password"] = "password needs to be at least 8 character."
	}
	return formErrors
}
