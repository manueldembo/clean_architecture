package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	List() ([]*Order, error)
	// GetTotal() (int, error)
}
