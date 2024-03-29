package main

import (
    "fmt"
    "html/template"
    "net/http"
)


func main() {
    http.HandleFunc("/", LoginPage)
    http.HandleFunc("/login", LoginPage)
    http.HandleFunc("/welcome", WelcomePage)


    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

   
    fmt.Println("Server started on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}


func SignupPage(w http.ResponseWriter, r *http.Request) {
 if r.Method == http.MethodPost {
 
  username := r.FormValue("username")
  password := r.FormValue("password")


  fmt.Printf("New user signup: Username - %s, Password - %s\n", username, password)

  http.Redirect(w, r, "/welcome", http.StatusSeeOther)
  return
 }


 tmpl, err := template.ParseFiles("templates/signup.html")
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 tmpl.Execute(w, nil)
}


func LoginPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        
        if username == "admin" && password == "admin" {
          
            http.Redirect(w, r, "/welcome", http.StatusSeeOther)
            return
        }

       
        fmt.Fprintf(w, "Invalid credentials. Please try again.")
        return
    }

    tmpl, err := template.ParseFiles("templates/login.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}


func WelcomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, you have successfully logged in!")
}