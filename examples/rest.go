package examples

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeloDev/go-server-template/database"
	"github.com/fatih/structs"
	_ "github.com/go-sql-driver/mysql"
)

type ExamplePost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        uint64 `json:"code"`
}

func (h *Handlers) Post(w http.ResponseWriter, r *http.Request) {

	// Only accept POST method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Validate request body
	var post ExamplePost
	bodyDecoder := json.NewDecoder(r.Body)
	bodyDecoder.DisallowUnknownFields()
	err := bodyDecoder.Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(post)

	// TODO: figure out required field validation during marshelling
	if post.Code == 0 || len(post.Title) < 1 || len(post.Description) < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Insert into db
	row, err := database.Insert("post", structs.Map(post))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Println(row)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Posted Successfully"))
}
