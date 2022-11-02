package request

type RequestDTO struct {
	FirstNumber   int    `json:"first_number"`
	SecondNumber  int    `json:"second_number"`
	NewPopulation int    `json:"new population"`
	CityId        int    `json:"city_id"`
	CityRegion    string `json:"city_region"`
	CityDistrict  string `json:"city_district"`
}
