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

func GetSellerTable(requestInfo s.RequestTableData) []s.SellerTableData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.SellerTableData

	//create Query
	var querys s.QueryInfo = query.CreateTableQuery(requestInfo)

	var sellerDataArr []s.SellerTableData
	for i := 0; i < len(querys.Querys); i++ {
		var query = querys.Querys[i]
		rows, err := db.Query(query.Query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var sellerData s.SellerTableData
		var data [7]int
		for rows.Next() {
			err := rows.Scan(&data[0], &data[1], &data[2], &data[3], &data[4], &data[5], &data[6])
			if err != nil {
				log.Fatal(err)
			}
		}
		sellerData.Header = query.Header
		sellerData.Id = i
		sellerData.TableData = data
		sellerDataArr = append(sellerDataArr, sellerData)
	}
	result = sellerDataArr
	return result
}

func GetOrderTable(requestInfo s.RequestTableData) []s.OrderTableData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.OrderTableData

	//create Query
	var querys s.QueryInfo = query.CreateTableQuery(requestInfo)
	var orderDataArr []s.OrderTableData
	for i := 0; i < len(querys.Querys); i++ {
		var query = querys.Querys[i]
		rows, err := db.Query(query.Query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var orderData s.OrderTableData
		var data [2]int
		for rows.Next() {
			err := rows.Scan(&data[0], &data[1])
			if err != nil {
				log.Fatal(err)
			}
		}
		orderData.Header = query.Header
		orderData.Id = i
		orderData.TableData = data
		orderDataArr = append(orderDataArr, orderData)
	}
	result = orderDataArr
	return result
}

func GetChart(requestInfo s.RequestChartData) s.ChartDataSet {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result s.ChartDataSet

	//create Query
	var query s.Querys = query.CreateChartQuery(requestInfo)
	var allDate []string = CreateAllDate(requestInfo)
	var dataSet s.ChartData
	rows, err := db.Query(query.Query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mockData s.MockChartData

	if requestInfo.Cycle == "daily" {
		for _, date := range allDate {
			mockData.Date = append(mockData.Date, date)
			mockData.Data = append(mockData.Data, 0)
		}
	}

	var number int
	var date string

	for rows.Next() {
		err := rows.Scan(&number, &date)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(date)
		date = strings.Replace(date, "-", ".", 2)
		date = strings.Replace(date, ".0", ".", 2)
		if requestInfo.Cycle == "daily" {
			for i, mockDate := range mockData.Date {
				if mockDate == date {
					mockData.Data[i] = number
				}
			}
		} else {
			mockData.Date = append(mockData.Date, date)
			mockData.Data = append(mockData.Data, number)
		}
	}
	dataSet.Data = mockData.Data
	dataSet.Name = requestInfo.Segment + " | " + requestInfo.Event
	result.Data = append(result.Data, dataSet)
	result.Categories = mockData.Date
	return result
}

func GetProducts() []s.ProductData {
	//db 연결
	db, err := sql.Open("mysql", db1)
	if err != nil {
		panic(err) //에러가 있으면 프로그램을 종료해라
	}
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var result []s.ProductData

	//create Query
	var query string = query.ProductsQuery

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

func CreateAllDate(requestInfo s.RequestChartData) []string {
	var result []string
	var startDate string = requestInfo.StartDate
	var endDate string = requestInfo.EndDate
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
