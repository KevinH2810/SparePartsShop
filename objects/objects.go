package objects

type (
	User struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		Password   string `json: "password"`
		Role       string `json:"role"`
		DateSignIn string `json"datesignin"`
	}

	Item struct {
		ID                  int    `json:"id"`
		Nama                string `json:"nama"`
		Code                string `json:"code"`
		Jenis               string `json:"jenis"`
		CompanyManufacturer string `json:"company"`
		BuyPrice            int    `json:"buyprice"`
		SellPrice1          int    `json:"sellprice1"`
		SellPrice2          int    `json:"sellprice2"`
	}
)
