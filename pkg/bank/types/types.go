package types

// Money for min (cent, cop, diram and other)
type Money int64 

//Currency is cod of currency
type Currency string

//cod of currency
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN in number of card
type PAN string

//Card info about card
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money // money use
	MinBalance Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
//payment is certificat about payments
type Payment struct {
	ID int
	Amount Money
}