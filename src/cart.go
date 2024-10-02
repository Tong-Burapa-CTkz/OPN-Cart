package cart

import (
	"fmt"
)

type Product struct {
	ID       string
	Price    float64
	Quantity int
}

type Discount struct {
	Name       string
	Amount     float64
	Percentage float64
	MaxAmount  float64
}

type Cart struct {
	Items     map[string]Product
	Discounts map[string]Discount
}

func NewCart() *Cart {
	return &Cart{
		Items:     make(map[string]Product),
		Discounts: make(map[string]Discount),
	}
}

func (c *Cart) AddProduct(productID string, price float64, quantity int) {
	if item, exists := c.Items[productID]; exists {
		item.Quantity += quantity
		c.Items[productID] = item
	} else {
		c.Items[productID] = Product{ID: productID, Price: price, Quantity: quantity}
	}
}

func (c *Cart) UpdateProduct(productID string, quantity int) {
	if item, exists := c.Items[productID]; exists {
		item.Quantity = quantity
		c.Items[productID] = item // Update with new value
	} else {
		fmt.Printf("Product ID %s not found in cart.\n", productID)
	}
}

func (c *Cart) RemoveProduct(productID string) {
	delete(c.Items, productID)
}

func (c *Cart) Destroy() {
	c.Items = make(map[string]Product)
	c.Discounts = make(map[string]Discount)
}

func (c *Cart) IsEmpty() bool {
	return len(c.Items) == 0
}

func (c *Cart) ListItems() []Product {
	items := make([]Product, 0, len(c.Items))
	for _, item := range c.Items {
		items = append(items, item)
	}
	return items
}

func (c *Cart) CountUniqueItems() int {
	return len(c.Items)
}

func (c *Cart) TotalItems() int {
	total := 0
	for _, item := range c.Items {
		total += item.Quantity
	}
	return total
}

func (c *Cart) TotalAmount() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func (c *Cart) ApplyDiscount(discount Discount) {
	c.Discounts[discount.Name] = discount
}

func (c *Cart) RemoveDiscount(discountName string) {
	delete(c.Discounts, discountName)
}

func (c *Cart) ApplyDiscounts() float64 {
	total := c.TotalAmount()

	for _, discount := range c.Discounts {
		if discount.Amount > 0 {
			total -= discount.Amount
		} else if discount.Percentage > 0 {
			discountAmount := total * discount.Percentage / 100
			if discountAmount > discount.MaxAmount {
				discountAmount = discount.MaxAmount
			}
			total -= discountAmount
		}
	}
	return total
}

func (c *Cart) ApplyFreebie(productID string, freebieID string) {
	if item, exists := c.Items[productID]; exists && item.Quantity > 0 {
		c.AddProduct(freebieID, 0, 1)
	}
}
