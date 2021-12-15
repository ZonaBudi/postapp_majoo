package domain

type ReportMerchant struct {
	Merchant     *PublicMerchant        `json:"merchant"`
	Transactions []*TransactionMerchant `json:"transactions"`
}

type ReportOutlet struct {
	Merchant     *PublicMerchant      `json:"merchant"`
	Outlet       *PublicOutlet        `json:"outlet"`
	Transactions []*TransactionOutlet `json:"transactions"`
}
