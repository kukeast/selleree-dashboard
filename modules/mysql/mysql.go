package mysql

import (
	"database/sql"
	"log"
	"main/modules/query"
	s "main/modules/structs"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var db1 string
var db2 string

func InitConnectionString(connectionStrings ...string) {
	db1 = connectionStrings[0]
	db2 = connectionStrings[1]
}

func GetProducts(limit string) []s.ProductData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.ProductData

	//create Query
	var query string = query.ProductsQuery
	query = query + limit
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data s.ProductData

	for rows.Next() {
		err := rows.Scan(&data.ItemName, &data.Price, &data.Visibility, &data.Deleted, &data.RawUrl, &data.RawCount, &data.StoreName, &data.ItemId, &data.StoreId)
		if err != nil {
			log.Fatal(err)
		}
		if data.RawUrl.Valid {
			data.Url = data.RawUrl.String
		} else {
			data.Url = ""
		}
		if data.RawCount.Valid {
			data.ImageCount = int(data.RawCount.Int32)
		} else {
			data.ImageCount = 0
		}
		result = append(result, data)
	}

	return result
}

func GetOrderDetail(orderId string) s.OrderDetailData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result s.OrderDetailData

	//create Query
	var query string = query.OrderDetailQuery
	query = query + orderId

	row := db.QueryRow(query)
	err = row.Scan(
		&result.BuyerName,
		&result.BuyerEmail,
		&result.BuyerCellPhoneNumber,
		&result.ZipCode,
		&result.AddressLine,
		&result.RawDetailLine,
		&result.RawBankName,
		&result.RawBankAccountNumber,
		&result.RawBankAccountHolder,
		&result.FinancialStatus,
		&result.FulfillmentStatus,
		&result.DefaultShippingFee,
		&result.ExtraShippingFee,
		&result.RawMemo,
		&result.CreatedAt,
		&result.LastModifiedAt,
		&result.PaymentMethod,
		&result.StoreName,
		&result.Identifier,
		&result.ItemName,
		&result.Price,
		&result.RawUrl,
		&result.Quantity,
		&result.ItemId)
	if err != nil {
		log.Fatal(err)
	}
	if result.RawUrl.Valid {
		result.ImageUrl = result.RawUrl.String
	} else {
		result.ImageUrl = ""
	}
	if result.RawBankName.Valid {
		result.BankName = result.RawBankName.String
	} else {
		result.BankName = ""
	}
	if result.RawBankAccountNumber.Valid {
		result.BankAccountNumber = result.RawBankAccountNumber.String
	} else {
		result.BankAccountNumber = ""
	}
	if result.RawBankAccountHolder.Valid {
		result.BankAccountHolder = result.RawBankAccountHolder.String
	} else {
		result.BankAccountHolder = ""
	}
	if result.RawMemo.Valid {
		result.Memo = result.RawMemo.String
	} else {
		result.Memo = ""
	}
	if result.RawDetailLine.Valid {
		result.AddressDetailLine = result.RawDetailLine.String
	} else {
		result.Memo = ""
	}
	return result
}

func GetOrders(limit string, sortBy string) []s.OrderData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.OrderData

	//create Query
	var query string = query.OrdersQuery
	query = query + " ORDER BY " + sortBy + " desc " + "LIMIT " + limit
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data s.OrderData

	for rows.Next() {
		err := rows.Scan(&data.OrderId, &data.OrderTitle, &data.CreatedAt, &data.LastModifiedAt, &data.DefaultShippingFee, &data.ExtraShippingFee, &data.Name, &data.Identifier, &data.Price, &data.Quantity, &data.RawUrl, &data.FinancialStatus, &data.FulfillmentStatus, &data.PaymentMethod)
		if err != nil {
			log.Fatal(err)
		}
		if data.RawUrl.Valid {
			data.ImageUrl = data.RawUrl.String
		} else {
			data.ImageUrl = ""
		}
		result = append(result, data)
	}

	return result
}

func GetShopggus() []s.ShopgguData {
	//db 연결
	db, err := sql.Open("mysql", db2)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.ShopgguData

	//create Query
	var query string = query.ShopggusQuery

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data s.ShopgguData

	for rows.Next() {
		err := rows.Scan(&data.StoreName, &data.Order, &data.Date)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, data)
	}

	return result
}

func GetToday() [][]s.TodayData {
	var results [][]s.TodayData

	//create Query
	for i := 0; i < 4; i++ {
		//db 연결
		var dbs string
		if i == 3 {
			dbs = db2
		} else {
			dbs = db1
		}
		db, err := sql.Open("mysql", dbs)

		if err != nil {
			panic(err) //에러가 있으면 프로그램을 종료해라
		}
		defer db.Close() //main함수가 끝나면 db를 닫아라
		var query string = query.TodayQuery[i]
		var result []s.TodayData
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var data s.TodayData

		for rows.Next() {
			err := rows.Scan(&data.Date, &data.Count)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, data)
		}
		results = append(results, result)
	}

	return results
}
func GetTodayChart(name string) s.ChartDataSet {
	//db 연결
	var dbs string
	if name == "published" {
		dbs = db2
	} else {
		dbs = db1
	}
	db, err := sql.Open("mysql", dbs)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result s.ChartDataSet

	//create Query
	var query string = query.TodayChartQuery[name]
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var currentTime = time.Now()
	var allDate = CreateAllDate(currentTime.AddDate(0, 0, -29).Format("2006-01-02"), currentTime.Format("2006-01-02"))
	var dataSet s.ChartData

	var mockData s.MockChartData
	for _, date := range allDate {
		mockData.Date = append(mockData.Date, date)
		mockData.Data = append(mockData.Data, 0)
	}

	var data int
	var date string

	for rows.Next() {
		err := rows.Scan(&date, &data)
		if err != nil {
			log.Fatal(err)
		}
		date = strings.Replace(date, "-", ".", 2)
		date = strings.Replace(date, ".0", ".", 2)
		for i, mockDate := range mockData.Date {
			if mockDate == date {
				mockData.Data[i] = data
			}
		}
		dataSet.Data = mockData.Data
	}
	if name == "store" {
		dataSet.Name = "상점 개설"
	} else if name == "item" {
		dataSet.Name = "상품 등록"
	} else if name == "order" {
		dataSet.Name = "주문"
	} else if name == "published" {
		dataSet.Name = "샵꾸 발행"
	}

	result.Data = append(result.Data, dataSet)
	result.Categories = mockData.Date

	return result
}

func GetFunnel(startDate string, endDate string) [7]string {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result [7]string

	//create Query
	var query string = query.FunnelQuery(startDate, endDate)

	row := db.QueryRow(query)
	err = row.Scan(
		&result[0],
		&result[1],
		&result[2],
		&result[3],
		&result[4],
		&result[5],
		&result[6],
	)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetPaymentSetting(startDate string, endDate string) [2][5]string {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result [2][5]string

	//create Query
	var query string = query.PaymentSettingQuery(startDate, endDate)

	row := db.QueryRow(query)
	err = row.Scan(
		&result[0][0],
		&result[0][1],
		&result[0][2],
		&result[0][3],
		&result[0][4],
		&result[1][0],
		&result[1][1],
		&result[1][2],
		&result[1][3],
		&result[1][4],
	)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetSellers(startDate string, endDate string, segment string, limit string) []s.SellerData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.SellerData
	//create Query
	var query string = query.SellersQuery(startDate, endDate, segment, limit)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data s.SellerData

	for rows.Next() {
		err := rows.Scan(&data.RawIdentifier, &data.RawName, &data.RawItemCount, &data.RawOrderCount)
		if err != nil {
			log.Fatal(err)
		}
		if data.RawIdentifier.Valid {
			data.Identifier = data.RawIdentifier.String
		} else {
			data.Identifier = ""
		}
		if data.RawName.Valid {
			data.Name = data.RawName.String
		} else {
			data.Name = ""
		}
		if data.RawItemCount.Valid {
			data.ItemCount = data.RawItemCount.String
		} else {
			data.ItemCount = ""
		}
		if data.RawOrderCount.Valid {
			data.OrderCount = data.RawOrderCount.String
		} else {
			data.OrderCount = ""
		}
		result = append(result, data)
	}

	return result
}

func CreateAllDate(startDate string, endDate string) []string {

	var result []string
	t, _ := time.Parse("2006-01-02", startDate)
	t2, _ := time.Parse("2006-01-02", endDate)
	days := int(t2.Sub(t).Hours() / 24)

	for i := 0; i < days+1; i++ {
		var dateString string = t.Format("2006.01.02")
		dateString = strings.Replace(dateString, ".0", ".", 2)
		result = append(result, dateString)
		t = t.AddDate(0, 0, 1)
	}
	return result
}
