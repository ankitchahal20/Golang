package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var noteStore = make(map[string]Note)

var id int = 0

/*
1. 	ResponseWriter is used for writing the response headers and bodies.
2. 	Header is written via writeHeader function of ResponseWriter
3. 	Response Body is wriiten with the help of the Write Method of ResponseWriter
*/

/*
1.	 All the information regarding the request is present in a pointer variable of type http.Request
2.	 Since this is json based API, so the request will be coming in the form of json.
3. 	 Hence to understand that, we first create a decode object, so that we can decode the request into the
	 json Object into our understandable format that is noteStore.
4. 	 The Vars function of the mux package returns the route variable for the current request.
*/

//HTTP POST - /api/notes/
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	//Decode the incoming json Note request
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	id++

	// convert the integer into string to store in the map.
	k := strconv.Itoa(id)
	//fmt.Println("K id :", k, " : ", id)
	noteStore[k] = note

	jsonObject, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonObject)
}

//HTTP GET - /api/notes/
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	//iterate through all the notes present in the map.
	for _, v := range noteStore {
		notes = append(notes, v)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonObject, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonObject)

}

//HTTP PUT - api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteTo Note
	//Decode the incoming request
	err = json.NewDecoder(r.Body).Decode(&noteTo)
	if err != nil {
		panic(err)
	}

	if _, ok := noteStore[k]; ok {
		//delete the existing item and add the update item
		delete(noteStore, k)
		noteStore[k] = noteTo
	} else {
		log.Printf("Unable to find the key %s in the Notes map to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	k := vars["id"]
	//fmt.Println("K : ", k)
	//Remove
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("Unable to find the key %s in the Notes map to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening...and enjoy")
	server.ListenAndServe()
}
