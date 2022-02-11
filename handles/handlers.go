package handles

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello home page")
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/all", getArticles).Methods("GET")
	r.HandleFunc("/one/{id}", checkArticle).Methods("GET")
	return r
}
