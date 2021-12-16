package query

var CoverQuery = func(id string) string {
	return `
		SELECT
			blocks.visible,
			covers.cover_media_url,
			covers.background_type,
			covers.background_color,
			covers.background_media_url
		FROM editor.malls as malls
		LEFT JOIN(
			SELECT *
			FROM editor.blocks
			WHERE type = 'COVER'
		) as blocks
		ON blocks.theme_id = malls.theme_id
		LEFT JOIN(
			SELECT *
			FROM editor.block_covers
		) as covers
		ON covers.block_id = blocks.id
		WHERE malls.seller_id = ` + id
}

var ProductsQuery = func(limit string, id string) string {
	var where string
	if id != "" {
		where = `Seller.id = ` + id
	} else {
		where = "store_id not in (1, 2, 9, 10, 49, 126, 209)"
	}
	return `
		SELECT Item.name, Item.price, Item.visibility,Item.deleted, Image.url, Image.c, Store.name, Item.id, Store.identifier
		FROM selleree.item AS Item
		LEFT JOIN(
			SELECT url, count(*) c, item_id
			FROM selleree.item_image
			GROUP BY item_id
		) AS Image
		ON Item.id = Image.item_id
		LEFT JOIN(
			SELECT id, name, identifier, seller_id
			FROM selleree.store
		) AS Store
		ON Store.id = Item.store_id
		LEFT JOIN(
			SELECT id
			FROM selleree.seller
		) AS Seller
		ON Store.seller_id = Seller.id
		WHERE ` + where + ` 
		ORDER BY Item.created_at desc
		LIMIT ` + limit
}
var OrdersQuery = func(limit string, sortBy string, id string) string {
	var where string
	if id != "" {
		where = `Seller.id = ` + id
	} else {
		where = "store_id not in (1, 2, 9, 10, 49, 126, 209)"
	}
	return `
		SELECT o.id, o.title, o.created_at, o.last_modified_at, o.default_shipping_fee, o.extra_shipping_fee, store.name, store.identifier, item.price, item.quantity, item.image_url, o.financial_status, o.fulfillment_status, o.payment_method
		FROM selleree.order AS o
		LEFT JOIN (
			SELECT *
			FROM selleree.store
		)AS store
		ON o.store_id = store.id
		LEFT JOIN (
			SELECT *
			FROM selleree.order_item
		)AS item
		ON item.order_id = o.id
		LEFT JOIN(
			SELECT id
			FROM selleree.seller
		) AS Seller
		ON store.seller_id = Seller.id
		WHERE ` + where + ` 
		ORDER BY ` + sortBy + ` desc LIMIT ` + limit
}
var ShopggusQuery string = `
	SELECT Mall.store_identifier, max(Mall_Theme.order), max(Mall_Theme.published_at)
	FROM editor.mall_themes AS Mall_Theme
	LEFT JOIN (
		SELECT *
		FROM editor.malls
	) AS Mall
	ON Mall.id = Mall_Theme.mall_id
	WHERE Mall.seller_id not in (1, 2, 3, 5, 55, 100, 149)
	GROUP BY Mall.store_identifier
	ORDER BY max(Mall_Theme.published_at) desc
	LIMIT 16
`

var OrderDetailQuery string = `
	SELECT o.buyer_name, o.buyer_email, o.buyer_cell_phone_number, o.zip_code, o.address_line, o.address_detail_line, o.bank_name, o.bank_account_number, o.bank_account_holder, o.financial_status, o.fulfillment_status, o.default_shipping_fee, o.extra_shipping_fee, o.memo, o.created_at, o.last_modified_at, o.payment_method, store.name, store.identifier, item.name, item.price, item.image_url, item.quantity, item.item_id, store.seller_id
	FROM selleree.order AS o
	LEFT JOIN (
		SELECT *
		FROM selleree.store
	)AS store
	ON o.store_id = store.id
	LEFT JOIN (
		SELECT *
		FROM selleree.order_item
	)AS item
	ON item.order_id = o.id
	WHERE o.id = 
`
var TodayQuery = [4]string{
	`
		SELECT date(created_at), count(created_at)
		FROM selleree.store
		WHERE DATE(created_at) = CURDATE() or DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)
		GROUP BY date(created_at)
		ORDER BY date(created_at) desc
	`,
	`
		SELECT date(created_at), count(created_at)
		FROM selleree.item
		WHERE (DATE(created_at) = CURDATE() or DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)) and store_id not in (1, 2, 9, 10, 49, 126, 209)
		GROUP BY date(created_at)
		ORDER BY date(created_at) desc
	`,
	`
		SELECT date(created_at), count(created_at)
		FROM selleree.order
		WHERE (DATE(created_at) = CURDATE() or DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)) and store_id not in (1, 2, 9, 10, 49, 126, 209)
		GROUP BY date(created_at)
		ORDER BY date(created_at) desc
	`,
	`
		SELECT date(themes.published_at) ,count(themes.published_at)
		FROM editor.mall_themes as themes
		left join(
			select * 
			from editor.malls
		) as malls
		on malls.id = themes.mall_id
		WHERE (DATE(themes.published_at) = CURDATE() or DATE(themes.published_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)) and malls.seller_id not in (1, 2, 3, 5, 55, 100, 149)
		GROUP BY date(themes.published_at)
		ORDER BY date(themes.published_at) desc
	`,
}
var TodayChartQuery = map[string]string{
	"store": `
		SELECT date(created_at) ,count(created_at)
		FROM selleree.store
		where date(created_at) > DATE_SUB(CURDATE(), INTERVAL 30 DAY)
		group by date(created_at)	
	`,
	"item": `
		SELECT date(created_at) ,count(created_at)
		FROM selleree.item
		where date(created_at) > DATE_SUB(CURDATE(), INTERVAL 30 DAY) and store_id not in (1, 2, 9, 10, 49, 126, 209)
		group by date(created_at)	
	`,
	"order": `
		SELECT date(created_at) ,count(created_at)
		FROM selleree.order
		where date(created_at) > DATE_SUB(CURDATE(), INTERVAL 30 DAY) and store_id not in (1, 2, 9, 10, 49, 126, 209)
		group by date(created_at)	
	`,
	"published": `
		SELECT date(themes.published_at) ,count(themes.published_at)
		FROM editor.mall_themes as themes
		left join(
			select * 
			from editor.malls
		) as malls
		on malls.id = themes.mall_id
		where date(themes.published_at) > DATE_SUB(CURDATE(), INTERVAL 30 DAY) and malls.seller_id not in (1, 2, 3, 5, 55, 100, 149)
		group by date(themes.published_at)	
	`,
}
var FunnelQuery = func(startDate string, endDate string) string {
	return `
		SELECT 
			sum(if(seller.id, 1, 0)) AS step1,
			sum(if(store.id, 1, 0)) AS step2,
			sum(if(JSON_LENGTH(payment.bank_accounts) or !payment.toss_contract ->> '$.contractStatus', 1, 0 )) AS step3,
			sum(if(item.itemCount >= 1 and payment.store_id, 1, 0 )) AS step4,
			sum(if(orders.orderCount >= 1, 1, 0)) AS step5,
			sum(if(orders.updatedCount >= 2, 1, 0)) AS step6,
			sum(if(orders.updatedCount >= 10, 1, 0)) AS step7
		FROM selleree.seller AS seller
		LEFT JOIN(
			SELECT seller_id, id
			FROM selleree.store
		) AS store
		ON store.seller_id = seller.id
		LEFT JOIN(
			SELECT store_id, count(id) AS itemCount
			FROM selleree.item
			WHERE DATE(created_at) <= DATE("` + endDate + `")
			GROUP BY store_id
		) AS item
		ON store.id = item.store_id
		LEFT JOIN(
			SELECT 
				created_at, 
				last_modified_at, 
				store_id, 
				count(id) AS orderCount,
				sum(if(last_modified_at != created_at, 1, 0)) AS updatedCount
			FROM selleree.order
			WHERE DATE(last_modified_at) <= DATE("` + endDate + `")
			GROUP BY store_id
		) AS orders
		ON store.id = orders.store_id
		LEFT JOIN(
			SELECT store_id, toss_contract, bank_accounts
			FROM selleree.payment_method
		) AS payment
		ON store.id = payment.store_id
		WHERE seller.id not in (1, 2, 3, 5, 55, 100, 149) and DATE(seller.created_at) >= DATE("` + startDate + `") and DATE(seller.created_at) <= DATE("` + endDate + `") 
	`
}

var PaymentSettingQuery = func(startDate string, endDate string) string {
	return `
		SELECT 
			sum(if(JSON_LENGTH(payment.bank_accounts) and toss_contract = CAST('null' AS JSON), 1, 0)) bank,
			sum(if(!payment.toss_contract ->> '$.contractStatus' and JSON_LENGTH(payment.bank_accounts) = 0, 1, 0)) toss,
			sum(if(JSON_LENGTH(payment.bank_accounts) and !payment.toss_contract ->> '$.contractStatus', 1, 0)) banktoss,
			sum(if(JSON_LENGTH(payment.bank_accounts) = 0 and toss_contract = CAST('null' AS JSON), 1, 0)) nothing,
			sum(if(payment.id, 1, 0)) payment_all,
			sum(if(payment.toss_contract ->> '$.contractStatus' = 'READY', 1, 0)) ready,
			sum(if(payment.toss_contract ->> '$.contractStatus' = 'WAIT_FOR_REVIEW', 1, 0)) waitforreview,
			sum(if(payment.toss_contract ->> '$.contractStatus' = 'DONE', 1, 0)) done,
			sum(if(payment.toss_contract ->> '$.contractStatus' = 'TERMINATED', 1, 0)) terminate,
			sum(if(!payment.toss_contract ->> '$.contractStatus', 1, 0)) toss_all
		FROM selleree.store AS store
		LEFT JOIN(
			SELECT *
			FROM selleree.payment_method
		) AS payment
		ON store.id = payment.store_id
		WHERE store.id not in (1, 2, 9, 10, 49, 126, 209) and DATE(store.created_at) >= DATE("` + startDate + `") and DATE(store.created_at) <= DATE("` + endDate + `") 
	`
}

var SellersQuery = func(startDate string, endDate string, segment string, limit string) string {
	var segments = map[string]string{
		"가입":              "seller.id",
		"상점 개설":           "store.id",
		"결제 설정":           "JSON_LENGTH(payment.bank_accounts) or !payment.toss_contract ->> '$.contractStatus'",
		"상품 1개 이상 등록":     "item.itemCount >= 1 and payment.store_id",
		"주문 1개 이상":        "orders.orderCount >= 1",
		"주문 상태 변경 2개 이상":  "orders.updatedCount >= 2",
		"주문 상태 변경 10개 이상": "orders.updatedCount >= 10",
		"무통장 입금":          "JSON_LENGTH(payment.bank_accounts) and toss_contract = CAST('null' AS JSON)",
		"토스페이먼츠":          "!payment.toss_contract ->> '$.contractStatus' and JSON_LENGTH(payment.bank_accounts) = 0",
		"무통장 입금 & 토스페이먼츠": "JSON_LENGTH(payment.bank_accounts) and !payment.toss_contract ->> '$.contractStatus'",
		"설정 안 함":          "JSON_LENGTH(payment.bank_accounts) = 0 and toss_contract = CAST('null' AS JSON)",
		"신청서 작성 중":        "payment.toss_contract ->> '$.contractStatus' = 'READY'",
		"신청 완료":           "payment.toss_contract ->> '$.contractStatus' = 'WAIT_FOR_REVIEW'",
		"심사 완료":           "payment.toss_contract ->> '$.contractStatus' = 'DONE'",
		"계약 종료":           "payment.toss_contract ->> '$.contractStatus' = 'TERMINATED'",
	}
	return `
		SELECT seller.id, seller.identifier, store.name, item.itemCount, orders.orderCount, store.company_information ->> '$.businessRegistrationNumber'
		FROM selleree.seller AS seller
		LEFT JOIN(
			SELECT seller_id, id, name, company_information
			FROM selleree.store
		) AS store
		ON store.seller_id = seller.id
		LEFT JOIN(
			SELECT store_id, count(id) AS itemCount
			FROM selleree.item
			WHERE DATE(created_at) <= DATE("` + endDate + `")
			GROUP BY store_id
		) AS item
		ON store.id = item.store_id
		LEFT JOIN(
			SELECT 
				created_at, 
				last_modified_at, 
				store_id, 
				count(id) AS orderCount,
				sum(if(last_modified_at != created_at, 1, 0)) AS updatedCount
			FROM selleree.order
			WHERE DATE(last_modified_at) <= DATE("` + endDate + `")
			GROUP BY store_id
		) AS orders
		ON store.id = orders.store_id
		LEFT JOIN(
			SELECT store_id, toss_contract, bank_accounts
			FROM selleree.payment_method
		) AS payment
		ON store.id = payment.store_id
		WHERE seller.id not in (1, 2, 3, 5, 55, 100, 149) and DATE(seller.created_at) >= DATE("` + startDate + `") and DATE(seller.created_at) <= DATE("` + endDate + `") and ` + segments[segment] + `
		ORDER BY seller.created_at desc
		LIMIT ` + limit
}

var SellerQuery = func(id string) string {
	return `
		SELECT 
			seller.id,
			seller.identifier, 
			seller.full_name,
			seller.cell_phone_number,
			seller.created_at,
			store.id,
			store.name,
			store.category,
			store.contacts,
			store.editor_used,
			store.design_published,
			item.itemCount,
			orders.orderCount,
			store.company_information ->> '$.businessRegistrationNumber',
			payment.bank_accounts ->> '$[0].holder',
			payment.bank_accounts ->> '$[0].bankName',
			payment.bank_accounts ->> '$[0].accountNumber',
			payment.toss_contract ->> '$.contractStatus'
		FROM selleree.seller AS seller
		LEFT JOIN(
			SELECT *
			FROM selleree.store
		) AS store
		ON store.seller_id = seller.id
		LEFT JOIN(
			SELECT store_id, count(id) AS itemCount
			FROM selleree.item
			GROUP BY store_id
		) AS item
		ON store.id = item.store_id
		LEFT JOIN(
			SELECT store_id, count(id) AS orderCount
			FROM selleree.order
			GROUP BY store_id
		) AS orders
		ON store.id = orders.store_id
		LEFT JOIN(
			SELECT store_id, toss_contract, bank_accounts
			FROM selleree.payment_method
		) AS payment
		ON store.id = payment.store_id
		WHERE seller.id = ` + id + `
		ORDER BY seller.created_at desc
	`
}
