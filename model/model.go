package model

//Reference Payload
type Reference struct {
	Amount   float64           `json:"amount"`
	DateEnd  string            `json:"end_datetime"`
	Customer map[string]string `json:"custom_fields"`
}

//Payment payment that were not yet Acknowledged by the client application ---Nullable	Max. Length	Description
type Payment struct {
	ID                    int               `json:"id"`                      // no	12Payment Id
	Amount                string            `json:"amount"`                  //	no	–	Payment amount
	Customfields          map[string]string `json:"custom_fields"`           //no	–	Key-value structure copied from the Reference
	Datetime              string            `json:"datetime"`                //	no	–	Payment datetime formatted as ISO8601
	EntityID              int               `json:"entity_id"`               //	no	5	Entity Id
	Fee                   string            `json:"fee"`                     //yes	–	Bank fee
	PeriodID              int               `json:"period_id"`               //	no	4	Period Id
	PeriodStartDatetime   string            `json:"period_start_datetime"`   //	no	–	Start datetime of the Period
	PeriodEndDatetime     string            `json:"period_end_datetime"`     //	no	–	End datetime of the Period
	ProductID             int               `json:"product_id"`              //	yes	2	Id of selected product (applicable if entity has multiple products)
	ParameterID           int               `json:"parameter_id"`            //	yes	2	Id of selected predefined amount (applicable if entity is configured to use predefined amounts)
	ReferenceID           int               `json:"reference_id"`            //	no	–	Reference Id that originated the Payment
	TransactionID         int               `json:"transaction_id"`          //	no	8	Unique transaction Id within the Period
	TerminalType          string            `json:"terminal_type"`           //	no	–	Type of terminal used for the payment (one of: ATM, POS, IB, MB)
	TerminalPeriodID      int               `json:"terminal_period_id"`      //	yes	3	Terminal period (only applicable when terminal_type is ATM or POS)
	TerminalTransactionID int               `json:"terminal_transaction_id"` //	yes	5	Terminal transaction (only applicable when terminal_type is ATM or POS)
	TerminalLocation      string            `json:"terminal_location"`       //	yes	15	Location as informed by the payment processor (only applicable when terminal_type is ATM or POS)
	TerminalID            string            `json:"terminal_id"`             //	yes	10	Id of the payment terminal (only applicable when terminal_type is ATM or POS)
}

//MockPayment Payload for mock payment
type MockPayment struct {
	ReferenceID int     `json:"reference_id"` //true	Reference Id
	Amount      float64 `json:"amount"`       //true	The payment amount
}
