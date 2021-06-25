package api

type Rovers struct {
	Rovers []Rover
}

type Rover struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string `json:"status"`
	MaxSol      int    `json:"max_sol"`
	MaxDate     string `json:"max_date"`
	TotalPhotos int    `json:"total_photos"`
	Cameras     []Camera
}

type Camera struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	RoverId  int    `json:"rover_id"`
	FullName string `json:"full_name"`
}
