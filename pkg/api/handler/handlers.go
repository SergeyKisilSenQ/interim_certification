package handlers

import (
	"context"
	"encoding/json"
	request "interim_certification/pkg/api/service"
	interfaces "interim_certification/pkg/storage"
	"interim_certification/pkg/storage/repo_cities"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetCity(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cityId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/city/"))
		cityInfo, _ := i.GetCity(context.TODO(), cityId)

		//Формирование ответа
		data, _ := json.Marshal(cityInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
func AddCity(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var c repo_cities.City
		if err := json.Unmarshal(content, &c); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		i.AddCity(context.TODO(), &c)

		//Формирование ответа
		data, _ := json.Marshal("Новый город зарегистрирован")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}
func DeleteCity(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var d request.RequestDTO
		if err := json.Unmarshal(content, &d); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//Формирование ответа
		var status int
		var data []byte
		if err := i.DeleteCity(context.TODO(), d.CityId); err != nil {
			data, _ = json.Marshal("Ошибка удаления")
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal("Город удален")
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}

func UpdateCityPopulation(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var p request.RequestDTO
		if err := json.Unmarshal(content, &p); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//Формирование ответа
		var status int
		var data []byte
		cityId, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/"))
		if err := i.UpdateCityPopulation(context.TODO(), cityId, p.NewPopulation); err != nil {
			data, _ = json.Marshal("Ошибка обновления данных")
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal("Информация обновлена")
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}
func GetRegion(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var a request.RequestDTO
		if err := json.Unmarshal(content, &a); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cityInfo := i.GetRegion(context.TODO(), a.CityRegion)

		//Формирование ответа
		data, _ := json.Marshal(cityInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
func GetDistrict(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var d request.RequestDTO
		if err := json.Unmarshal(content, &d); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cityInfo := i.GetDistrict(context.TODO(), d.CityDistrict)

		//Формирование ответа
		data, _ := json.Marshal(cityInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
func GetPopulation(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var p request.RequestDTO
		if err := json.Unmarshal(content, &p); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cityInfo := i.GetPopulation(context.TODO(), p.FirstNumber, p.SecondNumber)

		//Формирование ответа
		data, _ := json.Marshal(cityInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
func GetFoundation(i interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var f request.RequestDTO
		if err := json.Unmarshal(content, &f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cityInfo := i.GetFoundation(context.TODO(), f.FirstNumber, f.SecondNumber)

		//Формирование ответа
		data, _ := json.Marshal(cityInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
