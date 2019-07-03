package main

import "net/http"
import ("fmt";
		"log";
		"encoding/json")

func main() {
	port := 8080
	
	http.HandleFunc("/helloworld", helloWorldHandler)
	
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Frago: "Issuing WARNO for TPT 1562"}

	/*
	Receive and process a post response
	*/
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
   		http.Error(w, "Bad request", http.StatusBadRequest)
   		return
   	}


	/*
	Instead of using Marshal to deliver our data, we can simply use the
	built-in json encoder.
	*/
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	
	/*
	The code below was a convoluted way of taking some data and delivering it
	via the json.Marshall function.
	*/
	// data, err := json.Marshal(response)
	// if err!= nil {
	// 	panic("Oooops")
	// }

	// fmt.Fprint(w, string(data))

	// fmt.Fprint(w, "Hello World\n")
}

type helloWorldResponse struct {
	Frago string `json:"FRAGO"`
}