package main

type Player struct {
	Name	string	`json:"name"`
	Team	string	`json:"team"`
	Season	string	`json:"season"`
}

type Players []Player

func getFakeData() []Player {
	return []Player{
		Player{Name: "Connor McDavid", Team: "Edmonton Oilers", Season: "2016-2017"},
		Player{Name: "Patrick Kane", Team: "Chicago Blackhawks", Season: "2016-2017"},
	}
}
