package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "time"
  "github.com/gorilla/mux"
)


func main() {

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/404", notFound)
  router.HandleFunc("/400", badRequest)
  router.HandleFunc("/401", unAthorized )

  log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bukak sithik jhos!  %q", html.EscapeString(r.URL.Path))
}

func jamSekarang(w http.ResponseWriter, r *http.Request) {
  currentTime := time.Now().String()
  w.Write([]byte(currentTime))
}

func notFound(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusForbidden)
  fmt.Fprintf(w, "ampun mbukak niki lho, den. mboten pareng. http response: 404")
}

func badRequest(w http.ResponseWriter, r *http.Request ) {
  w.WriteHeader(http.StatusBadRequest)
  fmt.Fprintf(w, "njenenganipun napa sampun ngecek malih badhe mbikak napa?. http response: 400 ")
}

func unAthorized(w http.ResponseWriter, r *http.Request ) {
  w.WriteHeader(http.StatusUnauthorized)
  fmt.Fprintf(w, "njenengan kedah matur rumiyin sakderenge mbikak niki, nggih.  http reponse: 401 ")
}
