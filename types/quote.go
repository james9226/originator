package types

import "time"

type Quote struct {
	quoteID        string
	amount         int64
	term           int64
	apr            float64
	monthlyPayment float64
	totalCost      float64
	totalCharge    float64
}

type ApplicationOffer struct {
	applicationID       string
	customerID          string
	applicantHash       string
	requestedQuoteID	string 
	preApproved         bool
	approvalProbability int64
	applicationExpiry   time.Time
	quotes              []Quote
}
