package seeders

import (
	"retailStore/config"
	"retailStore/models"
)

func Seed() error {
	if _, err1 := ItemCategorySeed(); err1 != nil {
		return err1
	}
	if _, err2 := CouriersSeed(); err2 != nil {
		return err2
	}
	if _, err3 := PaymentServicesSeed(); err3 != nil {
		return err3
	}
	return nil
}

func ItemCategorySeed() ([]models.ItemCategory, error) {
	categoryNames := []string{"Electronic", "Sport", "Books"}
	id := uint(1)
	for _, categoryName := range categoryNames {
		category := models.ItemCategory{
			ID:           id,
			CategoryName: categoryName,
		}
		err := config.DB.Create(&category).Error
		id++
		if err != nil {
			return nil, err
		}
	}
	itemCategory := []models.ItemCategory{}
	config.DB.Find(&itemCategory)

	return itemCategory, nil

}
func CouriersSeed() ([]models.Courier, error) {
	companyNames := []string{"JNE", "J&T", "TIKI", "JET", "Wahana"}
	id := uint(1)
	for _, companyName := range companyNames {
		company := models.Courier{
			ID:          id,
			CompanyName: companyName,
		}
		err := config.DB.Create(&company).Error
		if err != nil {
			return nil, err
		}
		id++
	}
	couriers := []models.Courier{}
	//config.DB.Find(&couriers)

	return couriers, nil
}

func PaymentServicesSeed() ([]models.PaymentService, error) {
	paymentServices := map[string]string{
		"BCA m-Banking": "Sekarang semua transaksi perbankan #DibikinSimpel dengan BCA mobile",
		"OVO":           "From snack times to mealtimes, from routine bills to impulsive purchases, from online shopping to roadside stores - Pay everything and everywhere, with OVO!",
		"GO-PAY":        "Dompet digital yang memberikan promo terbaik!",
	}
	id := uint(1)
	for serviceName, description := range paymentServices {
		paymentService := models.PaymentService{
			ID:          id,
			CompanyName: serviceName,
			Description: description,
		}
		err := config.DB.Create(&paymentService).Error
		if err != nil {
			return nil, err
		}
		id++
	}
	paymentService := []models.PaymentService{}
	//config.DB.Find(&paymentService)

	return paymentService, nil

}

func ItemSeed() (models.ItemResponseArr, error) {
	name := []string{"naruto vol 1", "naruto vol 2", "mosquito racket", "solder", "mikasa volleyball", "molten volleyball"}
	category_id := []uint{3, 3, 1, 1, 2, 2}
	stock := []uint{20, 20, 10, 10, 30, 30}
	price := []uint{200, 200, 100, 100, 300, 300}
	id := uint(1)
	model := make([]models.Item, 6)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			temp := i*2 + j
			model[temp] = models.Item{
				ID:             id,
				Name:           name[temp],
				Description:    name[temp],
				Stock:          stock[temp],
				Price:          price[temp],
				ItemCategoryID: category_id[temp],
			}
			err := config.DB.Create(&model[temp]).Error
			if err != nil {
				return models.ItemResponseArr{}, err
			}
			id++
		}
	}
	modelResponse := []models.Item{}
	config.DB.Preload("ItemCategory").Find(&modelResponse)
	itemResponse := models.ItemResponseArr{
		Code:    200,
		Status:  "success",
		Message: "success getting items",
		Data:    modelResponse,
	}

	return itemResponse, nil
}

func UserSeed() (models.User, error) {
	user := models.User{
		ID:          1,
		Role:        "user",
		Password:    "kumenangismembayangkan",
		Name:        "Nurjamil",
		Email:       "nurjamil1996@gmail.com",
		Username:    "freferlay",
		PhoneNumber: "0895610234239",
	}
	// if err := config.DB.Create(&user).Error; err != nil {
	// 	return models.User{}, err
	// }
	// user.ShoppingCart.UserID = user.ID
	// user.ShoppingCart.ID = 1

	// if err := config.DB.Create(&user.ShoppingCart).Error; err != nil {
	// 	return models.User{}, err
	// }

	return user, nil
}
