package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {

	availableFormulas := map[string]interface{}{
		"availableFormulas": ConversorAvailableFormulas(),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(availableFormulas)
}

func ConverterHandler(res http.ResponseWriter, req *http.Request) {

	var convReq ConverterRequest

	body := json.NewDecoder(req.Body)
	defer req.Body.Close()

	if err := body.Decode(&convReq); err != nil {
		http.Error(res, "Invalid Request", http.StatusBadRequest)
		return
	}

	conv, err := NewConversor(convReq.From, convReq.Value, convReq.To)
	if err != nil {
		http.Error(res, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	result, err := conv.Result()
	if err != nil {
		http.Error(res, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	convRes := ConverterResponse{
		Value: result,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(convRes)
}
