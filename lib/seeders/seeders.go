package seeders

import (
	"fmt"
	"retailStore/config"
	"retailStore/models"
)

func Seed(){
	ItemCategorySeed()
	CouriersSeed()
	PaymentServicesSeed()
}

func ItemCategorySeed() {
	categoryNames := []string{"Electronic", "Books", "Sport"}
	for _, categoryName := range categoryNames {
		category := models.ItemCategory{
			CategoryName:categoryName,
		}
		err := config.DB.Save(&category).Error
	if err != nil {
		fmt.Println(err)
	}
	}
}
func CouriersSeed() {
	companyNames := []string{"JNE", "J&T", "TIKI", "JET", "Wahana"}

	for _, companyName := range companyNames {
		company := models.Courier{
			CompanyName:companyName,
		}
		err := config.DB.Save(&company).Error
		if err != nil {
		fmt.Println(err)
		}
	}
}

func PaymentServicesSeed() {
	paymentServices := map[string]string{
		"BCA m-Banking":"Sekarang semua transaksi perbankan #DibikinSimpel dengan BCA mobile",
		"OVO":"From snack times to mealtimes, from routine bills to impulsive purchases, from online shopping to roadside stores - Pay everything and everywhere, with OVO!",
		"GO-PAY":"Dompet digital yang memberikan promo terbaik!",
	}
	for serviceName,description := range paymentServices {
		paymentService := models.PaymentService{
			CompanyName:serviceName,
			Description: description,
		}
		err := config.DB.Save(&paymentService).Error
		if err != nil {
		fmt.Println(err)
		}
	}
	
}