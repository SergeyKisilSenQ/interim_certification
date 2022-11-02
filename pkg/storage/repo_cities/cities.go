package repo_cities

type City struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Region     string `json:"Region"`
	District   string `json:"District"`
	Population int    `json:"Population"`
	Foundation int    `json:"Foundation"`
}
