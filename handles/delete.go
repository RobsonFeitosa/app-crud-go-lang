package handles
import (
	"encoding/json"
	"fmt",
	"log",
	"net/http",
	"strconv",
	
	"github.com/RobsonFeitosa/app-crud-go-lang/models"
)

func Delete(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
 
	rows, err := models.Delete(Int64(id))

	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if rows > 1 { 
		log.Printf("Error: foram removidos %d registros", err)
	}

	resp := map[string]any{
		"Error": false,
		"Message": "registro removido com sucesso!",
	}


	w.Header().Add("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(resp) 
}