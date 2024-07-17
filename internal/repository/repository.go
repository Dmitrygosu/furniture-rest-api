package repository

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/Dmitrygosu/furniture-rest-api/internal/model"
)

// JSONDB - БД в формате JSON
type JSONDB struct {
	Path          string
	mu            sync.RWMutex
	FurnitureJSON FurnitureJSON
}

// FurnitureJSON - структура данных JSON файла
type FurnitureJSON struct {
	LastID         int64             `json:"last_id"`
	FurnitureArray []model.Furniture `json:"furniture_array"`
}

// NewJSONDB - инициализация новой JSON БД
func NewJSONDB(path string) JSONDB {
	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var furnJSON FurnitureJSON
	if err := json.Unmarshal(content, &furnJSON); err != nil {
		log.Println("База данных пустая, инициализация новой БД")
		return JSONDB{
			Path:          path,
			FurnitureJSON: FurnitureJSON{LastID: 0, FurnitureArray: []model.Furniture{}},
		}
	}

	return JSONDB{
		Path:          path,
		FurnitureJSON: furnJSON,
	}
}

// Create - добавление новой мебели в БД
func Create(data *model.Furniture, db *JSONDB) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.FurnitureJSON.LastID++
	data.ID = db.FurnitureJSON.LastID
	db.FurnitureJSON.FurnitureArray = append(db.FurnitureJSON.FurnitureArray, *data)

	return updateDB(db)
}

// Get - получение мебели по ID
func Get(id int64, db *JSONDB) *model.Furniture {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			return &furniture
		}
	}
	return nil
}

// Update - полное обновление
func Update(id int64, db *JSONDB, newFurniture *model.Furniture) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			newFurniture.ID = furniture.ID
			db.FurnitureJSON.FurnitureArray[i] = *newFurniture
			return updateDB(db)
		}
	}
	return nil
}

// Patch - частичное обновление
func Patch(id int64, db *JSONDB, newFurniture *model.Furniture) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			if newFurniture.Name != "" {
				db.FurnitureJSON.FurnitureArray[i].Name = newFurniture.Name
			}
			if newFurniture.Fabricator != "" {
				db.FurnitureJSON.FurnitureArray[i].Fabricator = newFurniture.Fabricator
			}
			if newFurniture.Height > 0 {
				db.FurnitureJSON.FurnitureArray[i].Height = newFurniture.Height
			}
			if newFurniture.Width > 0 {
				db.FurnitureJSON.FurnitureArray[i].Width = newFurniture.Width
			}
			if newFurniture.Length > 0 {
				db.FurnitureJSON.FurnitureArray[i].Length = newFurniture.Length
			}
			return updateDB(db)
		}
	}
	return nil
}

// Delete - удаление мебели по ID
func Delete(id int64, db *JSONDB) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			db.FurnitureJSON.FurnitureArray = append(db.FurnitureJSON.FurnitureArray[:i], db.FurnitureJSON.FurnitureArray[i+1:]...)
			return updateDB(db)
		}
	}
	return nil
}

// updateDB - обновление файла БД
func updateDB(db *JSONDB) error {
	data, err := json.Marshal(db.FurnitureJSON)
	if err != nil {
		return err
	}

	return os.WriteFile(db.Path, data, 0644)
}
