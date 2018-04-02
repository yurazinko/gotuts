package main

import ("fmt"
        "net/http")

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Expert web design by Yurii Zinko")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/about/", aboutHandler)
  http.ListenAndServe(":8000", nil)
}
