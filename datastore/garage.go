package datastore

import (
	"awesomeProject/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type garageHandler struct{}

func New() *garageHandler {
	return &garageHandler{}
}

func (g *garageHandler) GetAll(ctx *gofr.Context) ([]*models.Car, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT id, brand, model, owner, status, repairDate FROM garage")
	if err != nil {
		return nil, &errors.DB{Err: err}
	}
	defer rows.Close()

	cars := make([]*models.Car, 0)
	for rows.Next() {
		var c models.Car
		err = rows.Scan(&c.ID, &c.Brand, &c.Model, &c.Owner, &c.Status, &c.LastDateOfRepair)
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		cars = append(cars, &c)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	return cars, nil
}

func (g *garageHandler) Get(ctx *gofr.Context, id string) (*models.Car, error) {
	query := "SELECT id, brand, model, owner, status, repairDate FROM garage WHERE id=?"
	row := ctx.DB().QueryRowContext(ctx, query, id)

	var car models.Car
	err := row.Scan(&car.ID, &car.Brand, &car.Model, &car.Owner, &car.Status, &car.LastDateOfRepair)
	if err != nil {
		return nil, &errors.DB{Err: err}
	}
	return &car, nil
}

func (g *garageHandler) Add(ctx *gofr.Context, car *models.Car) (*models.Car, error) {
	query := "INSERT INTO garage (brand, model, owner, status, repairDate) VALUES (?, ?, ?, ?, ?)"
	res, err := ctx.DB().ExecContext(ctx, query, car.Brand, car.Model, car.Owner, car.Status, car.LastDateOfRepair)
	if err != nil {
		return nil, &errors.DB{Err: err}
	}
	id, _ := res.LastInsertId()
	car.ID = int(id)
	return car, nil
}

func (g *garageHandler) Update(ctx *gofr.Context, car *models.Car, id string) (*models.Car, error) {
	query := "UPDATE garage SET brand=?, model=?, owner=?, status=?, repairDate=? WHERE id=?"
	_, err := ctx.DB().ExecContext(ctx, query, car.Brand, car.Model, car.Owner, car.Status, car.LastDateOfRepair, id)
	if err != nil {
		return nil, &errors.DB{Err: err}
	}
	return car, nil
}

func (g *garageHandler) Delete(ctx *gofr.Context, id string) (*models.Car, error) {
	query := "DELETE FROM garage WHERE id=?"
	_, err := ctx.DB().ExecContext(ctx, query, id)
	if err != nil {
		return nil, &errors.DB{Err: err}
	}
	return nil, nil
}
