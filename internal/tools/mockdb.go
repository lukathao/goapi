package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"frank": {
		AuthToken: "ABC123",
		Username:  "frank",
	},
	"luka": {
		AuthToken: "EASYAS1",
		Username:  "luka",
	},
}

var mockCoinDetails = map[string]Coindetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"frank": {
		Coins:    400,
		Username: "frank",
	},
	"luka": {
		Coins:    1000,
		Username: "luka",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *Coindetails {
	time.Sleep(time.Second * 1)

	var clientData = Coindetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
