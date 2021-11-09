package query

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
	WHERE store_id not in (1, 2, 9, 10, 49, 126, 209)
	ORDER BY Item.created_at desc
	LIMIT 
`
var ShopggusQuery string = `
	SELECT Mall.store_identifier, Mall_Theme.order, Mall_Theme.published_at
	FROM editor.mall_themes AS Mall_Theme
	LEFT JOIN (
		SELECT *
		FROM editor.malls
	) AS Mall
	ON Mall.id = Mall_Theme.mall_id
	WHERE Mall.seller_id not in (1, 2, 3, 5, 55, 100, 149)
	ORDER BY Mall_Theme.published_at desc
	LIMIT 16
`

var OrdersQuery string = `
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
	WHERE store_id not in (1, 2, 9, 10, 49, 126, 209)
`
var OrderDetailQuery string = `
	SELECT o.buyer_name, o.buyer_email, o.buyer_cell_phone_number, o.zip_code, o.address_line, o.address_detail_line, o.bank_name, o.bank_account_number, o.bank_account_holder, o.financial_status, o.fulfillment_status, o.default_shipping_fee, o.extra_shipping_fee, o.memo, o.created_at, o.last_modified_at, o.payment_method, store.name, store.identifier, item.name, item.price, item.image_url, item.quantity, item.item_id
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
			sum(if(payment.store_id, 1, 0 )) AS step3,
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
			WHERE JSON_LENGTH(bank_accounts) != 0
			OR toss_contract -> '$.contractStatus' = 'WAIT_FOR_REVIEW'
			OR toss_contract -> '$.contractStatus' = 'DONE'
		) AS payment
		ON store.id = payment.store_id
		WHERE seller.id not in (1, 2, 3, 5, 55, 100, 149) and DATE(seller.created_at) >= DATE("` + startDate + `") and DATE(seller.created_at) <= DATE("` + endDate + `") 
	`
}

var FunnelDetailQuery = func(startDate string, endDate string, step string, limit string) string {
	var steps = map[string]string{
		"1": "seller.id",
		"2": "store.id",
		"3": "payment.store_id",
		"4": "item.itemCount >= 1 and payment.store_id",
		"5": "orders.orderCount >= 1",
		"6": "orders.updatedCount >= 2",
		"7": "orders.updatedCount >= 10",
	}
	return `
		SELECT seller.identifier, store.name, item.itemCount, orders.orderCount
		FROM selleree.seller AS seller
		LEFT JOIN(
			SELECT seller_id, id, name
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
			WHERE JSON_LENGTH(bank_accounts) != 0
			OR toss_contract -> '$.contractStatus' = 'WAIT_FOR_REVIEW'
			OR toss_contract -> '$.contractStatus' = 'DONE'
		) AS payment
		ON store.id = payment.store_id
		WHERE seller.id not in (1, 2, 3, 5, 55, 100, 149) and DATE(seller.created_at) >= DATE("` + startDate + `") and DATE(seller.created_at) <= DATE("` + endDate + `") and ` + steps[step] + `
		ORDER BY seller.created_at desc
		LIMIT ` + limit
}
