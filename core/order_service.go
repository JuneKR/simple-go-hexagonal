package core

type OrderService interface {
	CreateOrder(order Order) error
}
