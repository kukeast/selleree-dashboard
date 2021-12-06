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

type SellersData struct {
	Id                            string         `json:"id"`
	RawIdentifier                 sql.NullString `json:"-"`
	Identifier                    string         `json:"identifier"`
	RawName                       sql.NullString `json:"-"`
	Name                          string         `json:"name"`
	RawItemCount                  sql.NullString `json:"-"`
	ItemCount                     string         `json:"item_count"`
	RawOrderCount                 sql.NullString `json:"-"`
	OrderCount                    string         `json:"order_count"`
	BusinessRegistrationNumber    string         `json:"businessRegistrationNumber"`
	RawBusinessRegistrationNumber sql.NullString `json:"-"`
}

type SellerData struct {
	Id                         string `json:"id"`
	Identifier                 string `json:"identifier"`
	SellerName                 string `json:"seller_name"`
	PhoneNumber                string `json:"cell_phone_number"`
	CreatedAt                  string `json:"created_at"`
	StoreName                  string `json:"store_name"`
	Category                   string `json:"category"`
	Contacts                   string `json:"contacts"`
	EditorUsed                 string `json:"editor_used"`
	DesignPublished            string `json:"design_published"`
	ItemCount                  string `json:"item_count"`
	OrderCount                 string `json:"order_count"`
	BusinessRegistrationNumber string `json:"business_registration_number"`
	BankName                   string `json:"bank_name"`
	TossStatus                 string `json:"toss_status"`
}

type ShopgguData struct {
	StoreName string `json:"store_name"`
	Order     int    `json:"order"`
	Date      string `json:"date"`
}

type ProductData struct {
	ItemName   string         `json:"item_name"`
	ItemId     string         `json:"item_id"`
	Price      string         `json:"price"`
	Visibility string         `json:"visibility"`
	Deleted    string         `json:"deleted"`
	RawUrl     sql.NullString `json:"-"`
	Url        string         `json:"url"`
	RawCount   sql.NullInt32  `json:"-"`
	ImageCount int            `json:"image_count"`
	StoreName  string         `json:"store_name"`
	StoreId    string         `json:"store_id"`
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
