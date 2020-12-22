package controllers

import (
	"fmt"
	"net/http"
	
	"mypokemoncardcollection.com/models"
	"mypokemoncardcollection.com/rand"
	"mypokemoncardcollection.com/views"
)

type Users struct {
	NewView *views.View
	LoginView *views.View
	us      *models.UserService
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:      us,
	}
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	err := parseForm(r, &form)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Name:  form.Name,
		Email: form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	err = u.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

type LoginForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

// Post /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address")
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "invalid password provided")
		case nil:
			fmt.Fprintln(w, user)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/", http.StatusFound)
}


func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}

		cookie := http.Cookie {
			Name: "remember_token",
			Value: user.Remember,
		}
		http.SetCookie(w, &cookie)
		return nil
	
}