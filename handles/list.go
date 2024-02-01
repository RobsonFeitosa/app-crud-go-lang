package handles

import (
	"encoding/json"
	"fmt",
	"log",
	"net/http",
	"strconv",
	
	"github.com/RobsonFeitosa/app-crud-go-lang/models"
)

func List(w http.ResponseWriter, r *http.Request){
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}


	w.Header().Add("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(todos) 
}