package main

import "fmt"

// ExternalOrderService will call external service outside our SUT
type ExternalOrderService struct{}

func (exorder ExternalOrderService) DeductAmount(productName string, amountToDeduct int) error {
	fmt.Printf("[External Order Service] Deducting order for product %s, amount to deduct %v \n", productName, amountToDeduct)

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
	if amountSold < 1 {
		return fmt.Errorf("Wrong product amount, must be higher than 1")
	}

	if err := pos.orderService.DeductAmount(productName, amountSold); err != nil {
		return err
	}

	fmt.Printf("[POS Info] %s sold for %v piece \n", productName, amountSold)
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
