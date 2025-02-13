package bdd

import "time"

type Root struct {
	Films     string `json:"films"`
	People    string `json:"people"`
	Planets   string `json:"planets"`
	Species   string `json:"species"`
	Starships string `json:"starships"`
	Vehicles  string `json:"vehicles"`
}

type Starship struct {
	MGLT                 string        `json:"MGLT"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	CostInCredits        string        `json:"cost_in_credits"`
	Created              time.Time     `json:"created"`
	Crew                 string        `json:"crew"`
	Edited               time.Time     `json:"edited"`
	HyperdriveRating     string        `json:"hyperdrive_rating"`
	Length               string        `json:"length"`
	Manufacturer         string        `json:"manufacturer"`
	MaxAtmospheringSpeed string        `json:"max_atmosphering_speed"`
	Model                string        `json:"model"`
	Name                 string        `json:"name"`
	Passengers           string        `json:"passengers"`
	Films                []string      `json:"films"`
	Pilots               []interface{} `json:"pilots"`
	StarshipClass        string        `json:"starship_class"`
	Url                  string        `json:"url"`
}

type Film struct {
	Characters   []interface{} `json:"characters"`
	Created      time.Time     `json:"created"`
	Director     string        `json:"director"`
	Edited       time.Time     `json:"edited"`
	EpisodeId    int           `json:"episode_id"`
	OpeningCrawl string        `json:"opening_crawl"`
	Planets      []interface{} `json:"planets"`
	Producer     string        `json:"producer"`
	ReleaseDate  string        `json:"release_date"`
	Species      []interface{} `json:"species"`
	Starships    []interface{} `json:"starships"`
	Title        string        `json:"title"`
	Url          string        `json:"url"`
	Vehicles     []interface{} `json:"vehicles"`
}
