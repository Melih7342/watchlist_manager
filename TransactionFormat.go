package watchlist_manager

// ==========================================
// 1. MT Message Struct
// ==========================================
type MTMessage struct {
	TransactionRef string `json:"transactionReference"` // Field 20
	SenderBIC      string `json:"senderBic"`
	ReceiverBIC    string `json:"receiverBic"`

	OrderingCustomer    string `json:"orderingCustomer"`    // Field 50a/K (Zahler)
	BeneficiaryCustomer string `json:"beneficiaryCustomer"` // Field 59 (Empfänger)

	RemittanceInfo string `json:"remittanceInfo"` // Field 70
}

// ==========================================
// 2. MX Message Struct
// ==========================================
type MXMessage struct {
	MessageID string `json:"messageId"`

	// In MX ist alles viel feingranularer strukturiert
	DebtorName    string `json:"debtorName"`
	DebtorCountry string `json:"debtorCountry"`

	CreditorName    string `json:"creditorName"`
	CreditorCountry string `json:"creditorCountry"`

	RemittanceInfo string `json:"remittanceInfo"`
}

// ==========================================
// 3. Das Unified Screening Struct (WLM Input)
// ==========================================

type ScreeningRequest struct {
	TransactionID string
	FormatOrigin  string
	Entities      []TransactionEntity
	FreeText      string
}

type TransactionEntity struct {
	Role    string // e.g. "DEBTOR", "CREDITOR", "SENDER_BANK"
	Name    string
	Country string
}
