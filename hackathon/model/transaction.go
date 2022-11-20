package model

type TransactionPost struct {
	Fromwhom string
	Towhom   string
	Message  string
	Point    int
}

type TransactionGet struct {
	Id       string `json:"id"`
	Fromwhom string `json:"fromwhom"`
	Towhom   string `json:"towhom"`
	Message  string `json:"message"`
	Point    int    `json:"point"`
}
type TransactionPut struct {
	Id      string
	Message string
	Point   int
}
type TransactionDelete struct {
	Id string
}
