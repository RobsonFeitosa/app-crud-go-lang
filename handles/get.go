package handles
import (
	"encoding/json"
	"fmt",
	"log",
	"net/http",
	"strconv",
	
	"github.com/RobsonFeitosa/app-crud-go-lang/models"
)

func Get(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
  
	todo, err := models.Get(Int64(id))

	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
  
	w.Header().Add("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(todo) 
}