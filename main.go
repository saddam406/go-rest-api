package main
import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"

)
type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello home page")
}

func getArticles(w http.ResponseWriter,r *http.Request){
	articles:=Articles{
		Article{Title:"test title",Desc:"test desc",Content:"test content"},
		Article{Title:"sha title",Desc:"sha desc",Content:"sha content"},
	}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(articles)
}

func handleRequest(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/all",getArticles)
	log.Fatal(http.ListenAndServe(":9002",nil))
}
func main(){
	handleRequest()
}