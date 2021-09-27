package query

import (
	"main/modules/structs"
	"strings"
)

func CreateTableQuery(requestInfo structs.RequestTableData) structs.QueryInfo {
	var result structs.QueryInfo

	var dateRange string = CreateDateRange(requestInfo.StartDate, requestInfo.EndDate, requestInfo.Name)
	var tableQuerys = structs.NewQueryInfo()

	if requestInfo.Name == "seller" {
		for i := 0; i < len(sellerSegment); i++ {
			var segment = sellerSegment[i]
			var combination []string
			combination = append(combination, sellerFullQuery)
			combination = append(combination, segment["query"])
			combination = append(combination, dateRange)

			var query structs.Querys
			query.Header = segment["header"]
			query.Query = strings.Join(combination, " ")
			tableQuerys.Querys[i] = query

			result = tableQuerys
		}
	} else if requestInfo.Name == "order" {
		for i := 0; i < len(orderSegment); i++ {
			var segment = orderSegment[i]
			var combination []string
			combination = append(combination, orderFullQuery)
			combination = append(combination, segment["query"])
			combination = append(combination, dateRange)

			var query structs.Querys
			query.Header = segment["header"]
			query.Query = strings.Join(combination, " ")
			tableQuerys.Querys[i] = query

			result = tableQuerys
		}
	}
	return result
}

func CreateChartQuery(requestInfo structs.RequestChartData) structs.Querys {
	var result structs.Querys

	var dateRange string = CreateDateRange(requestInfo.StartDate, requestInfo.EndDate, requestInfo.Name)
	var dateCycle string = CreateDateCycle(requestInfo.Cycle, requestInfo.Name)
	var eventQuery string
	var segmentQuery string
	if requestInfo.Name == "seller" {
		for _, event := range sellerEvent {
			if requestInfo.Event == event["header"] {
				eventQuery = event["query"]
				break
			}
		}
		for _, segment := range sellerSegment {
			if requestInfo.Segment == segment["header"] {
				segmentQuery = segment["query"]
				break
			}
		}
		var combination []string
		combination = append(combination, eventQuery)
		combination = append(combination, ", ")
		combination = append(combination, dateCycle)
		combination = append(combination, sellerDefaultQuery)
		combination = append(combination, segmentQuery)
		combination = append(combination, dateRange)
		combination = append(combination, "GROUP BY ")
		combination = append(combination, dateCycle)

		var query structs.Querys
		query.Header = requestInfo.Segment
		query.Query = strings.Join(combination, " ")
		result = query
	} else if requestInfo.Name == "order" {
		for _, event := range orderEvent {
			if requestInfo.Event == event["header"] {
				eventQuery = event["query"]
				break
			}
		}
		for _, segment := range orderSegment {
			if requestInfo.Segment == segment["header"] {
				segmentQuery = segment["query"]
				break
			}
		}
		var combination []string
		combination = append(combination, eventQuery)
		combination = append(combination, ", ")
		combination = append(combination, dateCycle)
		combination = append(combination, orderDefaultQuery)
		combination = append(combination, segmentQuery)
		combination = append(combination, dateRange)
		combination = append(combination, "GROUP BY ")
		combination = append(combination, dateCycle)

		var query structs.Querys
		query.Header = requestInfo.Segment
		query.Query = strings.Join(combination, " ")
		result = query
	}
	return result
}

func CreateDateRange(startDate string, endDate string, name string) string {
	var dateRange []string
	var result string
	var dbTable string
	if name == "seller" {
		dbTable = "Seller"
	} else if name == "order" {
		dbTable = "Orders"
	}
	dateRange = append(dateRange, "date(")
	dateRange = append(dateRange, dbTable)
	dateRange = append(dateRange, ".created_at) >= date('")
	dateRange = append(dateRange, startDate)
	dateRange = append(dateRange, "')")
	if endDate != "" {
		dateRange = append(dateRange, "and date(")
		dateRange = append(dateRange, dbTable)
		dateRange = append(dateRange, ".created_at) <= date('")
		dateRange = append(dateRange, endDate)
		dateRange = append(dateRange, "')")
	}
	result = strings.Join(dateRange, " ")
	return result
}

func CreateDateCycle(cycle string, name string) string {
	var dateCycle string
	if name == "seller" {
		if cycle == "daily" {
			dateCycle = "date(Seller.created_at)"
		} else if cycle == "weekly" {
			dateCycle = "DATE_SUB(date(Seller.created_at), INTERVAL DAYOFWEEK(date(Seller.created_at)) - 1 DAY)"
		} else if cycle == "monthly" {
			dateCycle = "DATE_SUB(date(Seller.created_at), INTERVAL DAYOFMONTH(date(Seller.created_at)) - 1 DAY)"
		}
	} else if name == "order" {
		if cycle == "daily" {
			dateCycle = "date(Orders.created_at)"
		} else if cycle == "weekly" {
			dateCycle = "DATE_SUB(date(Orders.created_at), INTERVAL DAYOFWEEK(date(Orders.created_at)) - 1 DAY)"
		} else if cycle == "monthly" {
			dateCycle = "DATE_SUB(date(Orders.created_at), INTERVAL DAYOFMONTH(date(Orders.created_at)) - 1 DAY)"
		}
	}

	return dateCycle
}

var sellerDefaultQuery string = `
	FROM selleree.seller AS Seller
	LEFT JOIN (
			SELECT *
			FROM selleree.store
		)AS Store
	ON Store.seller_id = Seller.id
	LEFT JOIN (
			SELECT *
			FROM selleree.payment_method
		)AS Payment
	ON Payment.store_id = Store.id`

var sellerEvent = map[int]map[string]string{
	0: {
		"header": "계정 생성",
		"query":  "SELECT count(Seller.created_at)",
	},
	1: {
		"header": "본인 인증",
		"query":  "SELECT count(Seller.ci)",
	},
	2: {
		"header": "상점 생성",
		"query":  "SELECT count(Store.created_at)",
	},
	3: {
		"header": "계좌 등록",
		"query":  "SELECT count(Payment.type)",
	},
	4: {
		"header": "링크 등록",
		"query":  "SELECT count(Store.contacts)",
	},
	5: {
		"header": "샵꾸 연동",
		"query":  "SELECT count(if(Store.editor_used > 0, 1, null))",
	},
	6: {
		"header": "샵꾸 발행",
		"query":  "SELECT count(if(Store.design_published > 0, 1, null))",
	},
}

var sellerSegment = map[int]map[string]string{
	0: {
		"header": "전체",
		"query":  "WHERE ",
	},
	1: {
		"header": "남자",
		"query":  "WHERE Seller.sex = 'MAN' and",
	},
	2: {
		"header": "여자",
		"query":  "WHERE Seller.sex = 'WOMAN' and",
	},
	3: {
		"header": "20 ~ 24",
		"query":  "WHERE timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 > 19 and timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 < 25 and",
	},
	4: {
		"header": "25 ~ 29",
		"query":  "WHERE timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 > 24 and timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 < 30 and",
	},
	5: {
		"header": "30 ~ 34",
		"query":  "WHERE timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 > 29 and timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 < 35 and",
	},
	6: {
		"header": "35 ~ 39",
		"query":  "WHERE timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 > 34 and timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 < 40 and",
	},
	7: {
		"header": "40 ~",
		"query":  "WHERE timestampdiff(YEAR, Seller.birth_day, CURDATE()) + 1 > 39 and",
	},
	8: {
		"header": "패션",
		"query":  "WHERE Store.category = 'FASHION' and",
	},
	9: {
		"header": "뷰티",
		"query":  "WHERE Store.category = 'BEAUTY' and",
	},
	10: {
		"header": "팬시",
		"query":  "WHERE Store.category = 'FANCY_STATIONERY' and",
	},
	11: {
		"header": "푸드",
		"query":  "WHERE Store.category = 'FOOD' and",
	},
	12: {
		"header": "아트",
		"query":  "WHERE Store.category = 'ART' and",
	},
	13: {
		"header": "주얼리",
		"query":  "WHERE Store.category = 'JEWELRY_STUFF' and",
	},
	14: {
		"header": "건강 식품",
		"query":  "WHERE Store.category = 'HEALTH_FOOD' and",
	},
	15: {
		"header": "기타",
		"query":  "WHERE Store.category = 'ETC' and",
	},
}

var sellerFullQuery string = `
	SELECT 
		count(Seller.created_at), 
		count(Seller.ci), 
		count(Store.created_at), 
		count(Payment.type),
		count(Store.contacts), 
		count(if(Store.editor_used > 0, 1, null)),
		count(if(Store.design_published > 0, 1, null))
	FROM selleree.seller AS Seller
	LEFT JOIN (
		SELECT *
		FROM selleree.store
		)AS Store
	ON Store.seller_id = Seller.id
	LEFT JOIN (
		SELECT *
		FROM selleree.payment_method
		)AS Payment
	ON Payment.store_id = Store.id
`

//order
var orderDefaultQuery string = `
	FROM selleree.order as Orders
	LEFT JOIN (
		SELECT *
		FROM selleree.order_item
	)AS OrderItem
	ON OrderItem.order_id = Orders.id`

var orderEvent = map[int]map[string]string{
	0: {
		"header": "주문 수",
		"query":  "SELECT Count(Orders.created_at)",
	},
	1: {
		"header": "거래 금액",
		"query":  "SELECT sum(OrderItem.price)",
	},
}

var orderSegment = map[int]map[string]string{
	0: {
		"header": "전체",
		"query":  "WHERE ",
	},
	1: {
		"header": "결제 대기",
		"query":  "WHERE Orders.financial_status = 'WAITING' and",
	},
	2: {
		"header": "결제 완료",
		"query":  "WHERE Orders.financial_status = 'COMPLETE' and",
	},
	3: {
		"header": "주문 취소",
		"query":  "WHERE Orders.financial_status = 'CANCELED' and",
	},
	4: {
		"header": "배송 대기",
		"query":  "WHERE Orders.fulfillment_status = 'WAITING' and",
	},
	5: {
		"header": "배송 완료",
		"query":  "WHERE Orders.fulfillment_status = 'COMPLETE' and",
	},
	6: {
		"header": "배송 안 함",
		"query":  "WHERE Orders.fulfillment_status = 'WILL_NOT' and",
	},
}

var orderFullQuery string = `
	SELECT
		Count(Orders.created_at) as orders,
		sum(OrderItem.price) as price
	FROM selleree.order as Orders
	LEFT JOIN (
		SELECT *
		FROM selleree.order_item
	)AS OrderItem
	ON OrderItem.order_id = Orders.id`

var ProductsQuery string = `
	SELECT Item.name, Item.price, Item.visibility,Item.deleted, Image.url, Image.c, Store.name, Item.id, Store.identifier
	FROM selleree.item AS Item
	LEFT JOIN(
		SELECT url, count(*) c, item_id
		FROM selleree.item_image
		GROUP BY item_id
	) AS Image
	ON Item.id = Image.item_id
	LEFT JOIN(
		SELECT id, name, identifier
		FROM selleree.store
	) AS Store
	ON Store.id = Item.store_id
	ORDER BY Item.created_at desc
	LIMIT 
`
var ShopggusQuery string = `
	SELECT Mall.store_identifier, Mall_Theme.order, date_format(Mall_Theme.published_at, '%H시 %i분 %s초')
	FROM editor.mall_themes AS Mall_Theme
	LEFT JOIN (
		SELECT *
		FROM editor.malls
	) AS Mall
	ON Mall.id = Mall_Theme.mall_id
	ORDER BY Mall_Theme.published_at desc
	LIMIT 16
`
