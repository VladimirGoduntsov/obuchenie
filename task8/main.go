package main

import (
	"fmt"
	"task8/cars"
)

func main() {
	
	collection := &cars.AutoCollection{}

	newCar1 := cars.Car{Model: "Toyota", Year: 2025, Price: 50000.00}
	newCar2 := cars.Car{Model: "BMW", Year: 2018, Price: 37000.00}
	newCar3 := cars.Car{Model: "Audi", Year: 2021, Price: 48000.00}
	newCar4 := cars.Car{Model: "Hyundai", Year: 2011, Price: 15000.00}
	newCar5 := cars.Car{Model: "Kia", Year: 2020, Price: 31000.00}
	newCar6 := cars.Car{Model: "Skoda", Year: 2005, Price: 5000.00}

	collection.AddNewCar(newCar1)
	collection.AddNewCar(newCar2)
	collection.AddNewCar(newCar3)
	collection.AddNewCar(newCar4)
	collection.AddNewCar(newCar5)
	collection.AddNewCar(newCar6)

	
	collection.PrintInfo()

	
	fmt.Println("Sorted cars:")
	for _, car := range collection.GetSortedCars() {
		fmt.Printf("Modael: %s, Year: %d, Price: %.2f USD\n", car.Model, car.Year, car.Price)
	}

	
	collection.DeleteLastCar()

	
	fmt.Println("\nПосле удаления последней машины:")
	collection.PrintInfo()

	
	fmt.Println("Отсортированные машины:")
	for _, car := range collection.GetSortedCars() {
		fmt.Printf("Model: %s, Year: %d, Price: %.2f USD\n", car.Model, car.Year, car.Price)
	}

fmt.Println("Sorted cars:")
	for _, car := range collection.GetSortedCars2() {
		fmt.Printf("Model: %s, Year: %d, Price: %.2f USD\n", car.Model, car.Year, car.Price)
	}
}
	

