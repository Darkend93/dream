package repository

import (
	"context"
	"dream/entities"
	"dream/usecases/repository"
	"encoding/json"
	"io"
	"os"

	"gorm.io/gorm"
)

type dreamRepository struct {
	db *gorm.DB
}

func NewDreamRepository(db *gorm.DB) repository.DreamRepository {
	return &dreamRepository{db}
}

func (r *dreamRepository) Insert(c context.Context, dm *entities.Dream) error {
	return nil
}

func (r *dreamRepository) FindAll(c context.Context) ([]*entities.Dream, error) {
	var dreams entities.Dreams
	jsonFile, err := os.Open("db/dreamBase.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &dreams); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *dreamRepository) FindById(c context.Context, id *int) (*entities.Dream, error) {
	return nil, nil
}

func (r *dreamRepository) Update(c context.Context, de *entities.Dream) error {
	return nil
}
