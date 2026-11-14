package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"T-SHIRT": 20.00,
	"MUG":     12.50,
	"HAT":     18.00,
	"BOOK":    25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf("- Item %s (Sale! original: $%.2f, Sale Price: $%.2f)\n",
					originalItemCode, basePrice, salePrice)
			}
		}

		fmt.Printf("- Item: %s (Product not found)\n", itemCode)
		return 0.0, false
	}
	return basePrice, true
}

func main() {
	orderItems := []string{
		"T-SHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	var subtotal float64

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("subtotal Price: %.2f\n", subtotal)

}
