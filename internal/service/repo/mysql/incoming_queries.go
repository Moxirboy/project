package mysql

import (
	"fmt"
	"project/internal/configs"
	"project/internal/controller/v1/handler/adapter"
)

var (
	Db, _ = configs.DB()
)

func GetGoodsNames() []string {

	rows, err := Db.Query("select name from kirim.goods;")
	if err != nil {
		fmt.Println(64)
		panic(err)
	}
	defer rows.Close()
	var GoodsNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		GoodsNames = append(GoodsNames, name)
		if err != nil {
			fmt.Println(70)
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return GoodsNames
}

func InsertClientQuery() {
	query, err := Db.Query("INSERT INTO client (name, created) VALUES (?, ?)", adapter.DetailsClient.Name, adapter.DetailsClient.Date)
	defer query.Close()
	if err != nil {
		fmt.Println("InsertClientQuery")
		panic(err)
	}
}
func InsertProductQuery() {
	insertProductQuery := "INSERT INTO goods (name, sort) VALUES (?, ?)"
	_, err := Db.Query(insertProductQuery, adapter.Product.Name, adapter.Product.Sort)
	if err != nil {
		fmt.Println("InsertProductQuery")
		panic(err)
	}
}
func InsertPurchaseQuery() {
	clientID, err := GetProductIDByName(Db, "client", adapter.DetailsClient.Name)
	if err != nil {
		fmt.Println("clientId   InsertPurchaseQuery")
		panic(err)

	}

	insertPurchaseQuery := "INSERT INTO purchase (name, client_id) VALUES (?, ?)"
	_, err = Db.Query(insertPurchaseQuery, adapter.Product.Name, clientID)
	if err != nil {
		fmt.Println("from query   InsertPurchaseQuery")
		panic(err)
	}
}
func InsertPurchaseGoodsQuery() {
	productID, err := GetProductIDByName(Db, "goods", adapter.Product.Name)
	if err != nil {
		fmt.Println("productID   InsertPurchaseGoodsQuery")
		panic(err)
	}

	purchaseID, err := GetProductIDByName(Db, "purchase", adapter.Product.Name)
	if err != nil {
		fmt.Println("purchaseID   InsertPurchaseGoodsQuery")
		panic(err)
	}

	insertPurchaseGoodsQuery := "INSERT INTO purchase_goods (goods_id, purchase_id, amount,cort_price) VALUES (?, ?, ?,?)"
	_, err = Db.Query(insertPurchaseGoodsQuery, productID, purchaseID, adapter.Purchase.Amount, adapter.PurchaseGoods.CortPrice)
	if err != nil {
		fmt.Println("from query   InsertPurchaseGoodsQuery")
		panic(err)
	}
}
