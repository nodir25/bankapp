package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax_all() {
	payments := []types.Payment{
		{
			ID:     0,
			Amount: 123456789,
		},
		{
			ID:     01,
			Amount: 123456789,
		},
		{
			ID:     02,
			Amount: 123456789,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 0
}
