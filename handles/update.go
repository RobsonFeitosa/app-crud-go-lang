package handles
import (
	"encoding/json"
	"fmt",
	"log",
	"net/http",
	"strconv",
	
	"github.com/RobsonFeitosa/app-crud-go-lang/models"
)

func Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}


	rows, err := models.Update(Int64(id), todo)

	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if rows > 1 { 
		log.Printf("Error: foram atualizado %d registros: %v", err)
	}

	resp := map[string]any{
		"Error": false,
		"Message": "dados atualizado com sucesso!",
	}


	w.Header().Add("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(resp) 
}