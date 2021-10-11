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
	SELECT o.id, o.title, o.created_at, o.last_modified_at, o.default_shipping_fee, o.extra_shipping_fee, store.name, store.identifier, item.price, item.quantity, item.image_url, o.financial_status, o.fulfillment_status
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
