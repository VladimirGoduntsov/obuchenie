package cars

import (
	"fmt"
	"sort"
)

type Collection interface {
	AddNewCar(car Car) // Добавление машины 
	DeleteLastCar() // Удаление машины 
	SortCar() // Сортировка машин по году выпуска
	SortCar2() // Сортировка машин по прайсу
	PrintInfo() // Вывод информации о коллекции
	GetSortedCars() []Car // Получение отсортированных машин по году 
	GetSortedCars2() []Car // Получение отсортированных машин по прайсу
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

func (c *AutoCollection) SortCar() {          
	sort.Slice(c.cars, func(i, j int) bool {
		return c.cars[i].Year < c.cars[j].Year
	})
	fmt.Println("Сортировка по году выпуска:")
}

func (c *AutoCollection) SortCar2() {          
	sort.Slice(c.cars, func(i, j int) bool {
		return int(c.cars[i].Price) < int(c.cars[j].Price)
	})
	fmt.Println("Сортировка по прайсу:")
}


 func (c *AutoCollection) PrintInfo() {   
	fmt.Printf("Тип коллекции: ArrayCollection, Количество Элементов: %d\n", len(c.cars) )
 }

 func (c *AutoCollection) GetSortedCars() []Car {  
	c.SortCar()
	return c.cars
 }
 
 func (c *AutoCollection) GetSortedCars2() []Car {  
	c.SortCar2()
	return c.cars
 }
 
