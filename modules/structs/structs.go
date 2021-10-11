package structs

import "database/sql"

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TodayData struct {
	Date  string `json:"date"`
	Count string `json:"count"`
}

type ShopgguData struct {
	StoreName string `json:"store-name"`
	Order     int    `json:"order"`
	Date      string `json:"date"`
}

type ProductData struct {
	ItemName   string         `json:"item-name"`
	ItemId     string         `json:"item-id"`
	Price      string         `json:"price"`
	Visibility string         `json:"visibility"`
	Deleted    string         `json:"deleted"`
	RawUrl     sql.NullString `json:"-"`
	Url        string         `json:"url"`
	RawCount   sql.NullInt32  `json:"-"`
	ImageCount int            `json:"image-count"`
	StoreName  string         `json:"store-name"`
	StoreId    string         `json:"store-id"`
}

type OrderData struct {
	OrderId            string         `json:"id"`
	OrderTitle         string         `json:"title"`
	CreatedAt          string         `json:"created_at"`
	LastModifiedAt     string         `json:"last_modified_at"`
	DefaultShippingFee string         `json:"default_shipping_fee"`
	ExtraShippingFee   string         `json:"extra_shipping_fee"`
	Name               string         `json:"name"`
	Identifier         string         `json:"identifier"`
	Price              string         `json:"price"`
	Quantity           string         `json:"quantity"`
	ImageUrl           string         `json:"image_url"`
	RawUrl             sql.NullString `json:"-"`
	FinancialStatus    string         `json:"financial_status"`
	FulfillmentStatus  string         `json:"fulfillment_status"`
}

type ChartDataSet struct {
	Categories []string    `json:"categories"`
	Data       []ChartData `json:"data"`
}

type ChartData struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

type MockChartData struct {
	Data []int
	Date []string
}
