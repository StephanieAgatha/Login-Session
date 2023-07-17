package controller

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"learn-session-login-logout/config"
	"learn-session-login-logout/entities"
	"learn-session-login-logout/models"
	"net/http"
)

var userModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", 303)
		return
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", 303)
			return
		} else {
			welcome := map[string]interface{}{
				"FullName": session.Values["FullName"],
			}
			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, welcome)
		}
	}
}

type LoginInput struct {
	Username string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		//logic login
		r.ParseForm()
		userinput := &LoginInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		var user entities.User
		userModel.Where(&user, "username", userinput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("Username Not Found !")
		} else {
			hashedpw := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userinput.Password))
			if hashedpw != nil {
				message = errors.New("Wrong password !")
			}
		}
		//jika login gagal,maka tampilkan
		if message != nil {
			data := map[string]interface{}{
				"Error": message,
			}
			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {
			//set cookie
			session, _ := config.Store.Get(r, config.SESSION_ID)
			//set session value
			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["password"] = user.Password
			session.Values["FullName"] = user.FullName

			//save session
			session.Save(r, w)
			http.Redirect(w, r, "/", 303)
		}
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", 303)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		user := entities.User{
			FullName:          r.Form.Get("FullName"),
			Email:             r.Form.Get("email"),
			Username:          r.Form.Get("username"),
			Password:          r.Form.Get("password"),
			ConfirmedPassword: r.Form.Get("confirmpassword"),
		}
		var message = make(map[string]interface{})
		if user.FullName == "" {
			message["FullName"] = "Full Name cannot be empty"
		}
		if user.Email == "" {
			message["email"] = "Email cannot be empty"
		}
		if user.Username == "" {
			message["username"] = "Username cannot be empty"
		}
		if user.Password == "" {
			message["password"] = "Password cannot be empty"
		} else {
			if user.ConfirmedPassword != user.Password {
				message["configmpassword"] = "Password didn't match"
			}
		}

		if (len(message)) > 0 {
			data := map[string]interface{}{
				"Error": message,
			}
			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		} else {
			hashpass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashpass)
		}
		//insert to database
		_, err := userModel.Create(user)

		var loginmessage string
		if err != nil {
			loginmessage = "Register Failed" + loginmessage
		} else {
			loginmessage = "Successfully Register !"
		}
		data := map[string]interface{}{
			"Message": loginmessage,
		}
		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, data)
	}
}
