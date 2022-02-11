package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

var articles = Articles{
	Article{ID: "1", Title: "test title", Desc: "test desc", Content: "test content"},
	Article{ID: "2", Title: "sha title", Desc: "sha desc", Content: "sha content"},
}

func getArticles(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
func checkArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	get := vars["id"]
	fmt.Fprintf(w, get+"  this article is present\n")
	for _, value := range articles {
		if value.ID == get {
			json.NewEncoder(w).Encode(value)
		}
	}

}
