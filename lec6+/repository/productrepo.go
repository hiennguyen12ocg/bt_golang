package repository

import (
	"bt/project/connect"
	"bt/project/models"
	"fmt"
)

var db = connect.Connect()

func InitData() {
	defer db.Close()
	listProducts := []models.Product{
		{
			Name:        "A",
			Description: "Description A",
			Price:       340,
			CategoryId:  1,
			Image:       "Hahaha",
			Sale:        30,
		},
		{
			Name:        "B",
			Description: "Description B",
			Price:       40,
			CategoryId:  2,
			Image:       "Ahihi",
			Sale:        10,
		},
	}

	for _, product := range listProducts {
		strQuery := "INSERT INTO products(name, description, price, category_id, image, sale) VALUES (?,?,?,?,?,?)"
		result, err := db.Exec(strQuery, product.Name, product.Description, product.Price, product.CategoryId, product.Image, product.Sale)
		if err != nil {
			panic(err.Error())
		}
		lastID, _ := result.LastInsertId()
		fmt.Printf("Sản phẩm mới thêm có ID là: %d\n", lastID)
	}
}

func GetAllProducts() (result []models.Product) {

	kq, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	var product models.Product
	for kq.Next() {
		err := kq.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Image, &product.Sale)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%v\n", product)
		result = append(result, product)
	}
	defer kq.Close()
	return result

}
