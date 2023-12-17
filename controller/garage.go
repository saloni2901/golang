package controller

import (
	"awesomeProject/datastore"
	"awesomeProject/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type garageHandler struct {
	store datastore.GarageStore
}

func New(store datastore.GarageStore) *garageHandler {
	return &garageHandler{store: store}
}

func (g *garageHandler) GetAll(ctx *gofr.Context) (interface{}, error) {
	res, err := g.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *garageHandler) Get(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")
	res, err := g.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *garageHandler) Add(ctx *gofr.Context) (interface{}, error) {
	var car models.Car
	if err := ctx.Bind(&car); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	res, err := g.store.Add(ctx, &car)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *garageHandler) Update(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")

	var car models.Car
	if err := ctx.Bind(&car); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	res, err := g.store.Update(ctx, &car, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *garageHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")
	res, err := g.store.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
