package main

import "testing"

func TestPointOfSales_ProductSold(t *testing.T) {
	type fields struct {
		orderService OrderService
	}
	type args struct {
		productName string
		amountSold  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := PointOfSales{
				orderService: tt.fields.orderService,
			}
			if err := pos.ProductSold(tt.args.productName, tt.args.amountSold); (err != nil) != tt.wantErr {
				t.Errorf("PointOfSales.ProductSold() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
