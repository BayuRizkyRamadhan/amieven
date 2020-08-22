package main

import "fmt"

// ExternalOrderService will call external service outside our SUT
type ExternalOrderService struct{}

func (exorder ExternalOrderService) DeductAmount(productName string, amountToDeduct int) error {

	return nil
}

// OrderService will deduct amount if any transaction happens
type OrderService interface {
	DeductAmount(string, int) error
}

// Implement our order service
type PointOfSales struct {
	orderService OrderService
}

func (pos PointOfSales) ProductSold(productName string, amountSold int) error {

	return nil
}

func main() {
	fmt.Println("Welcome to the POS")

	externalOrderService := ExternalOrderService{}
	posService := PointOfSales{externalOrderService}

	err := posService.ProductSold("Baygon", 1000)
	if err != nil {
		fmt.Println(err)
	}
}
