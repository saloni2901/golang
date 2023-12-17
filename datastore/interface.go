package datastore

import (
	"awesomeProject/models"
	"gofr.dev/pkg/gofr"
)

type GarageStore interface {
	GetAll(ctx *gofr.Context) ([]*models.Car, error)
	Get(ctx *gofr.Context, id string) (*models.Car, error)
	Add(ctx *gofr.Context, car *models.Car) (*models.Car, error)
	Update(ctx *gofr.Context, car *models.Car, id string) (*models.Car, error)
	Delete(ctx *gofr.Context, id string) (*models.Car, error)
}
