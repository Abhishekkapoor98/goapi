package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDeatils{
	"Abhishek": {
		AuthToken: "123ABC",
		Username:  "Abhishek",
	},
	"Gaurav": {
		AuthToken: "456DEF",
		Username:  "Gaurav",
	},
	"Manish": {
		AuthToken: "789GHI",
		Username:  "Manish",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Abhishek": {
		Coins:    1000,
		Username: "Abhishek",
	},
	"Gaurav": {
		Coins:    1500,
		Username: "Gaurav",
	},
	"Manish": {
		Coins:    2000,
		Username: "Manish",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDeatils {

	time.Sleep(50 * time.Millisecond) // Simulate DB latency
	var clientData = LoginDeatils{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (m *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(50 * time.Millisecond) // Simulate DB latency
	var coinData = CoinDetails{}
	coinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &coinData
}

func (m *mockDB) SetupDatabase() error {
	return nil
}
