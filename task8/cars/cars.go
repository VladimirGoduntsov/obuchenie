package cars

import (
	"fmt"
	"sort"
)

type Collection interface {
	AddNewCar(car Car) // Добавление машины 
	DeleteLastCar() // Удаление машины 
	SortCarByYear() // Сортировка машин по году выпуска
	SortCarByPrice() // Сортировка машин по прайсу
	PrintInfo() // Вывод информации о коллекции
	GetCars() []Car // Получение слайс машин  
}


type Car struct {
	Model string
	Year int
	Price float64
}


type AutoCollection struct {
	cars []Car
}

func (c *AutoCollection) AddNewCar(car Car) {  
	c.cars = append(c.cars, car)
}

func (c *AutoCollection) DeleteLastCar() {  
	if len(c.cars) > 0 {
		c.cars = c.cars[:len(c.cars)-1]
	}
}      
	// SortCarByYear - этот метод выполняет сортировку коллекции по году 
func (c *AutoCollection) SortCarByYear() {          
	sort.Slice(c.cars, func(i, j int) bool {
		return c.cars[i].Year < c.cars[j].Year
	})	
}
// SortCarByYear - этот метод выполняет сортировку коллекции по прайсу 
func (c *AutoCollection) SortCarByPrice() {          
	sort.Slice(c.cars, func(i, j int) bool {
		return int(c.cars[i].Price) < int(c.cars[j].Price)
	})
}


 func (c *AutoCollection) PrintInfo() {   
	fmt.Printf("Тип коллекции: AutoCollection, Количество Элементов: %d\n", len(c.cars) )
 }

 func (c *AutoCollection) GetCars() []Car {  
	return c.cars
 }
 

 
