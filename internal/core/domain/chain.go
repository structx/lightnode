package domain

// SigWallet
type SigWallet struct{}

// Wallet
type Wallet struct{}

// Chain
type Chain interface{}

// Block
type Block struct {
	Hash          string        `json:"hash"`
	PrevHash      string        `json:"prev_hash"`
	Timestamp     string        `json:"timestamp"`
	Data          []byte        `json:"data"`
	Transactions  []Transaction `json:"transactions"`
	AccessCtrlRef string        `json:"access_ctrl_ref"`
	AccessHash    string        `json:"access_hash"`
}

// Transaction
type Transaction struct {
	ID            []byte   `json:"id"`
	Type          string   `json:"type"`
	Sender        string   `json:"sender"`
	Receiver      string   `json:"receiver"`
	Data          []byte   `json:"data"`
	Signatures    []string `json:"signatures"`
	AccessCtrlRef string   `json:"access_ctrl_ref"`
}
