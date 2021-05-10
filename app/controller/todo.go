package controller

import (
	"html/template"
	"net/http"
	"swa__prakt2_todo-02/app/model"
	"time"
)

var tmpl *template.Template

// Is executed automatically on package load
// Reads in and parses HTML templates
func init() {
	tmpl = template.Must(template.ParseGlob("app/view/*.html"))
}

// Index controller
// If logged-in, the todo list is shown, otherwise the login page
func Index(w http.ResponseWriter, r *http.Request) {
	sid, _ := r.Cookie("sid")

	if sid != nil && sid.Value != "" && model.Session[sid.Value] != "" {
		var tmplData = struct {
			Username string
		}{
			Username: model.Session[sid.Value],
		}
		tmpl.ExecuteTemplate(w, "todo.html", tmplData)
	} else {
		tmpl.ExecuteTemplate(w, "login.html", nil)
	}
}

// Login controller
// Checks credentials and sets session if credentials are OK
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username != "" && model.User[username] == password {
		// generate random session id
		sid, _ := generateRandomString(32)

		// add entry to session store
		model.Session[sid] = username

		// write session id to cookie
		// Expires 	-> Ablauf nach einem Jahr
		// Path 	-> Zugriff nur bei Paths die mit "/" beginnen
		// HttpOnly -> XSS (Cross Site Scripting) -> kein JavaScript Zugriff auf Cookies
		// SameSite	-> Cookie wird nur von aktueller (first) Domain verwaltet und nicht durch third-party Domains.
		//				Bsp. blog.seite wird auf evil.example referenziert und man will nicht,
		//				dass evil.example der Cookie von blog.seite mitgeschickt wird
		cookie := http.Cookie{Name: "sid",
			Value:    sid,
			HttpOnly: true,
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			Path:     "/",
			SameSite: http.SameSiteStrictMode}
		http.SetCookie(w, &cookie)

		var tmplData = struct {
			Username string
		}{
			Username: username,
		}
		tmpl.ExecuteTemplate(w, "todo.html", tmplData)
	} else {
		var tmplData = struct {
			ErrorMessage string
		}{
			ErrorMessage: "Username and/or password wrong!",
		}
		tmpl.ExecuteTemplate(w, "login.html", tmplData)
	}
}

// Logout controller
// Logs user out
func Logout(w http.ResponseWriter, r *http.Request) {
	sid, _ := r.Cookie("sid")
	delete(model.Session, sid.Value)

	cookie := http.Cookie{Name: "sid", Value: ""}
	http.SetCookie(w, &cookie)

	tmpl.ExecuteTemplate(w, "login.html", nil)
}
