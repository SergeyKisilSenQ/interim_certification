package repo_func

import (
	"context"
	"encoding/csv"
	"errors"
	"interim_certification/pkg/storage/repo_cities"
	"os"
	"strconv"
)

type CityStorage map[int]*repo_cities.City

func (Cs CityStorage) Put(Id *repo_cities.City) {
	Cs[Id.Id] = Id
}
func (Cs CityStorage) GetCity(cxt context.Context, cityId int) (*repo_cities.City, error) {
	c, ok := Cs[cityId]
	if !ok {
		return nil, errors.New("Такого города не существует")
	} else {
		return c, nil
	}
}
func (Cs CityStorage) GetRegion(ctx context.Context, cityRegion string) []*repo_cities.City {
	city := make([]*repo_cities.City, 0)
	if len(Cs) > 0 {
		for i, c := range Cs {
			if Cs[i].Region == cityRegion {
				city = append(city, c)
			}
		}
	}
	return city
}
func (Cs CityStorage) GetDistrict(ctx context.Context, cityDistrict string) []*repo_cities.City {
	city := make([]*repo_cities.City, 0)
	if len(Cs) > 0 {
		for i, c := range Cs {
			if Cs[i].District == cityDistrict {
				city = append(city, c)
			}
		}
	}
	return city
}

func (Cs CityStorage) GetPopulation(ctx context.Context, FirstNumber, SecondNumber int) []*repo_cities.City {
	city := make([]*repo_cities.City, 0)
	if len(Cs) > 0 {
		for i, c := range Cs {
			if Cs[i].Population >= FirstNumber && Cs[i].Population <= SecondNumber {
				city = append(city, c)
			}
		}
	}
	return city
}
func (Cs CityStorage) GetFoundation(ctx context.Context, FirstNumber, SecondNumber int) []*repo_cities.City {
	city := make([]*repo_cities.City, 0)
	if len(Cs) > 0 {
		for i, c := range Cs {
			if Cs[i].Foundation >= FirstNumber && Cs[i].Foundation <= SecondNumber {
				city = append(city, c)
			}
		}
	}
	return city
}
func (Cs CityStorage) AddCity(ctx context.Context, city *repo_cities.City) {
	Cs[city.Id] = city
}
func (Cs CityStorage) DeleteCity(ctx context.Context, cityId int) error {
	if _, err := Cs.GetCity(ctx, cityId); err != nil {
		return err
	}
	delete(Cs, cityId)
	return nil
}

func (Cs CityStorage) UpdateCityPopulation(ctx context.Context, cityId, population int) error {
	if _, err := Cs.GetCity(ctx, cityId); err != nil {
		return err
	}
	Cs[cityId].Population = population
	return nil
}

func NewCityStorage() CityStorage {
	return make(map[int]*repo_cities.City)
}

func ReadFile(cs CityStorage) (Cs CityStorage) {
	file, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		IdCity, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		NameCity := record[1]
		RegionCity := record[2]
		DistrictCity := record[3]
		PopulationCity, err := strconv.Atoi(record[4])
		if err != nil {
			panic(err)
		}
		FoundationCity, err := strconv.Atoi(record[5])
		if err != nil {
			panic(err)
		}
		NewCity := repo_cities.City{
			Id:         IdCity,
			Name:       NameCity,
			Region:     RegionCity,
			District:   DistrictCity,
			Population: PopulationCity,
			Foundation: FoundationCity,
		}
		cs.Put(&NewCity)
	}
	return Cs
}
