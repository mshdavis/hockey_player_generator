package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    // Pull out header setting stuff into helpers eventually
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(getFakeData()); err != nil {
        panic(err)
    }
}

func PredictPlayer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "It's Patrick Kane (you wish)!")
}

func PlayerInfo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    playerId:= vars["playerId"]
    fmt.Fprintln(w, "Player Info:", playerId)
}