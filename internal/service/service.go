package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Dmitrygosu/furniture-rest-api/internal/model"
	"github.com/Dmitrygosu/furniture-rest-api/internal/repository"
	"github.com/gorilla/mux"
)

// Service предоставляет методы для работы с мебелью
type Service struct {
	JSONDB repository.JSONDB
}

// NewService создает новый Service
func NewService(path string) *Service {
	return &Service{
		JSONDB: repository.NewJSONDB(path),
	}
}

// Create обрабатывает создание новой сущности мебели
func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	req := new(model.Furniture)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if req.HasEmptyFields() {
		responseError(w, http.StatusBadRequest, errors.New("Not enough fields to create entity | Не хватает полей для создания сущности"))
		return
	}

	if err := repository.Create(req, &s.JSONDB); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		log.Printf("Error while adding data to db: %s | Ошибка при добавлении данных в БД: %s", err, err)
		return
	}

	response(w, http.StatusCreated, req)
}

// Get обрабатывает получение сущности мебели по ID
func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	furniture := repository.Get(int64(id), &s.JSONDB)
	if furniture == nil {
		responseError(w, http.StatusNotFound, errors.New("Furniture not found | Мебель не найдена"))
		return
	}

	response(w, http.StatusOK, furniture)
}

// GetAll обрабатывает получение всех сущностей мебели
func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	furnitureArray := s.JSONDB.FurnitureJSON.FurnitureArray
	if len(furnitureArray) == 0 {
		response(w, http.StatusNoContent, nil)
		return
	}
	response(w, http.StatusOK, furnitureArray)
}

// Update обрабатывает обновление сущности мебели
func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(model.Furniture)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if req.HasEmptyFields() {
		responseError(w, http.StatusBadRequest, errors.New("Not enough fields to update entity | Не хватает полей для обновления сущности"))
		return
	}

	if err := repository.Update(int64(id), &s.JSONDB, req); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, req)
}

// Patch обрабатывает частичное обновление сущности мебели
func (s *Service) Patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(model.Furniture)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := repository.Patch(int64(id), &s.JSONDB, req); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

// Delete обрабатывает удаление сущности мебели
func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.Delete(int64(id), &s.JSONDB); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func response(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println(err)
		}
	}
}

func responseError(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error": err.Error()})
}
