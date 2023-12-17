package main

import (
	"awesomeProject/controller"
	"awesomeProject/datastore"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	d := datastore.New()

	c := controller.New(d)

	app.GET("/car/:id", c.Get)
	app.GET("/cars", c.GetAll)
	app.POST("/car", c.Add)
	app.PUT("/car/:id", c.Update)
	app.DELETE("/car/:id", c.Delete)
}
