package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

// FUNC RETURN INSTANCE OF ORDER SERVICE
func newOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

// BUSINESS LOGIC
func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
