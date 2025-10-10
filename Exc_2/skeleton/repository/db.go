package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	// drinks := ...
	drinks := []model.Drink{
		{ID: 1, Name: "the John", Description: "Big blue", Price: 90},
		{ID: 2, Name: "the Jane", Description: "Big pink", Price: 10},
		{ID: 3, Name: "the Rog", Description: "Big white", Price: 0},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Now(), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Now(), Amount: 3},
		{DrinkID: 3, CreatedAt: time.Now(), Amount: 4},
		{DrinkID: 3, CreatedAt: time.Now(), Amount: 8},
		{DrinkID: 1, CreatedAt: time.Now(), Amount: 8},

	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}


func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	var totalledOrders = map[uint64]uint64{}
	for _, v := range db.orders{
		_, ok := totalledOrders[v.DrinkID]
		if ok{
			totalledOrders[v.DrinkID] += uint64(v.Amount)
		} else{
			totalledOrders[v.DrinkID] = uint64(v.Amount)
		}
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
