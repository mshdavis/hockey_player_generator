package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        Name: "Index",
        Method: "GET",
        Pattern: "/",
        HandlerFunc: Index,
    },
    Route{
        Name: "PredictPlayer",
        Method: "GET",
        Pattern: "/predict",
        HandlerFunc: PredictPlayer,
    },
    Route{
        Name: "PlayerInfo",
        Method: "GET",
        Pattern: "/player/{playerId}",
        HandlerFunc: PlayerInfo,
    },
}