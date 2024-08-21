package auth

import (
	"errors"
	"fmt"
	"net/http"
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
		_, err = h.Store.CreateUser(ctx, store.CreateUserParams{
			Name:     name,
			Email:    email,
			Password: password,
		})
		// Checking for creation errors
		if err != nil {
			fmt.Println("error creating user.")
			return
		}
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
		if user.Password == password {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}
		fmt.Println("incorrect password")
	})
}

func checkForCredentials(name string, email string, pass string) map[string]string {
	formErrors := make(map[string]string)
	if len(name) < 3 {
		formErrors["name"] = "name needs to be more than 2 characters."
	}
	if len(email) < 3 || !strings.Contains(email, "@") || !strings.Contains(email[strings.Index(email, "@"):], ".") {
		formErrors["name"] = "email needs to be valid (e.g. example@gmail.com)."
	}
	if len(pass) < 8 {
		formErrors["password"] = "password needs to be at least 8 character."
	}
	return formErrors
}
