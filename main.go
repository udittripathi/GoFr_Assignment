package main

import (
	"gofr.dev/pkg/gofr"
	"encoding/json"
	"github.com/udittripathi/GoFr_Assignment/models"
)


func main() {
	// initialise gofr object
	app := gofr.New()


	app.POST("/carinfo", func(ctx *gofr.Context) (interface{}, error) {
		//var carReq Car
        var carReq models.Car
		// Parse JSON request body
		if err := json.NewDecoder(ctx.Request().Body).Decode(&carReq); err != nil {
			return nil, err
		}

		// Inserting a customer row in database using SQL
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO cars (make, model, entry_time, repair_status) VALUES (?, ?, ?, ?)", carReq.Make, carReq.Model, carReq.EntryTime, carReq.RepairStatus)
	

		return nil, err
	})

	app.GET("/car", func(ctx *gofr.Context) (interface{}, error) {

		var Cars []models.Car

		// Getting the Car from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Cars")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var Car models.Car
			if err := rows.Scan(&Car.ID, &Car.Make, &Car.Model, &Car.EntryTime, &Car.RepairStatus); err != nil {
				return nil, err
			}

			Cars = append(Cars, Car)
		}

		// return the customer
		return Cars, nil
	})

	app.GET("/car/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Get the car ID from the path parameters
		carID := ctx.PathParam("id")
	
		// Query the database for a specific car by ID
		row := ctx.DB().QueryRowContext(ctx, "SELECT * FROM Cars WHERE ID = ?", carID)
	
		var car models.Car
		// Scan the result into the car struct
		if err := row.Scan(&car.ID, &car.Make, &car.Model, &car.EntryTime, &car.RepairStatus); err != nil {
			return nil, err
		}
	
		// Return the single car
		return car, nil
	})

	app.POST("/completeRepair/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Get the car ID from the path parameters
		carID := ctx.PathParam("id")
	
		// Update the repair status in the database
		_, err := ctx.DB().ExecContext(ctx, "UPDATE cars SET repair_status = 'Completed' WHERE id = ?", carID)
		if err != nil {
			return nil, err
		}
	
		// Return success or any desired response
		return "Repair status updated to Completed successfully", nil
	})


	app.DELETE("/deleteCar/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Get the car ID from the path parameters
      carID := ctx.PathParam("id")

    // Delete the car entry from the database
    _, err := ctx.DB().ExecContext(ctx, "DELETE FROM cars WHERE id = ?", carID)
    if err != nil {
        return nil, err
    }

    // Return a response indicating success
    return map[string]interface{}{
        "message": "Car entry deleted successfully",
    }, nil
	})


	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}