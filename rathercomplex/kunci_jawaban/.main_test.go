package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type externalOrderMock struct {
	mock.Mock
}

// mock our external order service
func (eom *externalOrderMock) DeductAmount(productName string, amountToDeduct int) error {
	fmt.Printf("mock called with value %s , %v \n", productName, amountToDeduct)

	return nil
}

func TestPointOfSales_ProductSold(t *testing.T) {
	// initialize mock on beginning of test
	eoService := new(externalOrderMock)

	// create point of sales service
	posService := PointOfSales{eoService}

	type fields struct {
		pos PointOfSales
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
		{
			name: "Test fill in the product ",
			args: args{
				productName: "baygon",
				amountSold:  10,
			},
			fields: fields{
				pos: posService,
			},
			wantErr: false,
		},
		{
			name: "Test failed fill in the product because amount < 1 ",
			args: args{
				productName: "baygon",
				amountSold:  0,
			},
			fields: fields{
				pos: posService,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.pos.ProductSold(tt.args.productName, tt.args.amountSold); (err != nil) != tt.wantErr {
				t.Errorf("PointOfSales.ProductSold() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
