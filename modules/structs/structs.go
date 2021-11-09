package structs

import "database/sql"

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

type TodayData struct {
	Date  string `json:"date"`
	Count string `json:"count"`
}

type FunnelDetail struct {
	RawIdentifier sql.NullString `json:"-"`
	Identifier    string         `json:"identifier"`
	RawName       sql.NullString `json:"-"`
	Name          string         `json:"name"`
	RawItemCount  sql.NullString `json:"-"`
	ItemCount     string         `json:"itemCount"`
	RawOrderCount sql.NullString `json:"-"`
	OrderCount    string         `json:"orderCount"`
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
	PaymentMethod      string         `json:"payment_method"`
}

type OrderDetailData struct {
	BuyerName            string         `json:"buyer_name"`
	BuyerEmail           string         `json:"buyer_email"`
	BuyerCellPhoneNumber string         `json:"buyer_cell_phone_number"`
	ZipCode              string         `json:"zip_code"`
	AddressLine          string         `json:"address_line"`
	AddressDetailLine    string         `json:"address_detail_line"`
	RawDetailLine        sql.NullString `json:"-"`
	BankName             string         `json:"bank_name"`
	RawBankName          sql.NullString `json:"-"`
	BankAccountNumber    string         `json:"bank_account_number"`
	RawBankAccountNumber sql.NullString `json:"-"`
	BankAccountHolder    string         `json:"bank_account_holder"`
	RawBankAccountHolder sql.NullString `json:"-"`
	Memo                 string         `json:"memo"`
	RawMemo              sql.NullString `json:"-"`
	DefaultShippingFee   string         `json:"default_shipping_fee"`
	ExtraShippingFee     string         `json:"extra_shipping_fee"`
	FinancialStatus      string         `json:"financial_status"`
	FulfillmentStatus    string         `json:"fulfillment_status"`
	CreatedAt            string         `json:"created_at"`
	LastModifiedAt       string         `json:"last_modified_at"`
	PaymentMethod        string         `json:"payment_method"`
	StoreName            string         `json:"store_name"`
	Identifier           string         `json:"identifier"`
	ItemName             string         `json:"item_name"`
	Price                string         `json:"price"`
	Quantity             string         `json:"quantity"`
	ImageUrl             string         `json:"image_url"`
	RawUrl               sql.NullString `json:"-"`
	ItemId               string         `json:"item_id"`
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
