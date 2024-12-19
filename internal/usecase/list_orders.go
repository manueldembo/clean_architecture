package usecase

import "github.com/manueldembo/clean_architecture/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	dto := make([]OrderOutputDTO, len(orders))
	for i, order := range orders {
		dto[i] = OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return dto, nil
}
