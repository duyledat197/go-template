package wallet

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/stamp-server/config"
	"github.com/stamp-server/models"
)

type Service interface {
	Create(userID string) (string, string, error)
}

type service struct {
	walletRepo models.WalletRepository
	userRepo   models.UserRepository
}

func (s *service) Create(userID string) (string, string, error) {
	_, err := s.userRepo.FindByUserID(userID)
	if err != nil {
		return "", "", err
	}
	url := "http://api-exchange.cse30.io"
	var jsonStr = []byte(
		` {
					"jsonrpc":"2.0",
					"method":"cse_createWallet",
					"params":[""],
					"id":1
				}
			`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", config.APIKey)
	req.Header.Set("apisecret", config.APISecret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return "", "", err
	}
	defer resp.Body.Close()
	var result CSEAddress
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", "", err
	}

	address := result.Result.Address
	privateKey := result.Result.Pk.PrivateKey
	uuidNew, err := uuid.NewUUID()
	wallet := models.Wallet{
		ID:        uuid.UUID.String(uuidNew),
		UserID:    userID,
		PublicKey: privateKey,
		Address:   address,
	}

	err = s.walletRepo.Create(&wallet)
	if err != nil {
		return "", "", err
	}

	return address, privateKey, nil
}

// NewService ...
func NewService(walletRepo models.WalletRepository, userRepo models.UserRepository) Service {
	return &service{
		walletRepo: walletRepo,
		userRepo:   userRepo,
	}
}

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
