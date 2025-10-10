package model

type Drink struct {
	ID uint64 `json:"id"`
	// Add fields: Name, Price, Description with suitable types
	// json attributes need to be snakecase, i.e. name, created_at, my_variable, .
	Name string `json:"name"`
	Description string `json:"description"`
	Price uint64`json:"price"`
	
}
