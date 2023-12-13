package main

import (
	"gofr.dev/pkg/gofr"
	"encoding/json"
)

type Car struct {
	ID          string    `json:"id"`
	Make        string    `json:"make"`
	Model       string    `json:"model"`
	EntryTime   string `json:"entry_time"`
	RepairStatus string   `json:"repair_status"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		// Get the value using the redis instance
		value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()

        return value, err
    })

	app.POST("/carinfo", func(ctx *gofr.Context) (interface{}, error) {
		var carReq Car

		// Parse JSON request body
		if err := json.NewDecoder(ctx.Request().Body).Decode(&carReq); err != nil {
			return nil, err
		}

		// Inserting a customer row in database using SQL
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO cars (make, model, entry_time, repair_status) VALUES (?, ?, ?, ?)", carReq.Make, carReq.Model, carReq.EntryTime, carReq.RepairStatus)
	

		return nil, err
	})

	// app.GET("/car", func(ctx *gofr.Context) (interface{}, error) {
	// 	var Cars []Car

	// 	// Getting the Car from the database using SQL
	// 	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Cars")
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	for rows.Next() {
	// 		var Car Car
	// 		if err := rows.Scan(&Car.ID, &Car.Name); err != nil {
	// 			return nil, err
	// 		}

	// 		Cars = append(Cars, Car)
	// 	}

	// 	// return the customer
	// 	return Cars, nil
	// })

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}