package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	maximalPayment := payments [0]
	for _, payment := range payments {
		if maximalPayment.Amount < payment.Amount {
			maximalPayment = payment
		}
	}

	return maximalPayment
}
