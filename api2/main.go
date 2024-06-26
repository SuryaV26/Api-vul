package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var flag = "flag{AP2-Broken_Authentication}"

func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("login").Parse(`
        <html>
        <body>
            <form method="POST" action="/login">
                Username: <input type="text" name="username"><br>
                Password: <input type="password" name="password"><br>
                <input type="submit" value="Login">
            </form>
        </body>
        </html>
    `)
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if pass, ok := users[username]; ok && pass == password {
		fmt.Fprintf(w, "Welcome %s! Here is your flag: %s", username, flag)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", login)

	fmt.Println("Server starting on :8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
