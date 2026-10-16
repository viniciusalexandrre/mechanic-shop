package customer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	var customer CustomerRequest

	if err = json.Unmarshal(bodyRequest, &customer); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(customer)
}
