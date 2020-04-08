package models

// Privatekey ...
type Privatekey struct {
	PrivateKey string `json:"privateKey"`
	Iv         string `json:"iv"`
	Salt       string `json:"salt"`
}

// PrivatekeyRs ...
type PrivatekeyRs struct {
	Pk      Privatekey `json:"pk"`
	Address string     `json:"address"`
}

// CSEAddress ...
type CSEAddress struct {
	Jsonrpc string       `json:"jsonrpc" `
	ID      int          `json:"id"`
	Result  PrivatekeyRs `json:"result"`
}
