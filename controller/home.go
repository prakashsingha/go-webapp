package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/prakashsingha/go-webapp/model"
	"github.com/prakashsingha/go-webapp/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.homeHandler)
	http.HandleFunc("/login", h.loginHandler)
}

func (h home) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}

func (h home) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error logging in: %s\n", err.Error())
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if user, err := model.Login(email, password); err == nil {
			log.Printf("User has logged in: %v", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("Cannot login: %v", err)
			vm.Email = email
			vm.Password = password
		}
	}

	h.loginTemplate.Execute(w, vm)
}
