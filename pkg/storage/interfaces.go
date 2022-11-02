package interfaces

import (
	"context"
	"interim_certification/pkg/storage/repo_cities"
)

type Repository interface {
	GetCity(ctx context.Context, cityId int) (*repo_cities.City, error)
	AddCity(ctx context.Context, city *repo_cities.City)
	DeleteCity(ctx context.Context, cityId int) error
	UpdateCityPopulation(ctx context.Context, cityId, population int) error
	GetRegion(ctx context.Context, cityRegion string) []*repo_cities.City
	GetDistrict(ctx context.Context, cityDistrict string) []*repo_cities.City
	GetPopulation(ctx context.Context, FirstNumber, SecondNumber int) []*repo_cities.City
	GetFoundation(ctx context.Context, FirstNumber, SecondNumber int) []*repo_cities.City
}
