package main

import "fmt"

func main() {
	PaySuccessTemplate := "Hi! New Purchase Made in %s. \nUserId: %s\nPremium Type: %s\nBrand: %s %s\n"
	text = fmt.Sprintf(PaySuccessTemplate,
		u.Country,
		sub.UserId,
		sub.GadgetPremiumPlan.ProductPlan.PlanType.StringVal,
		sub.GadgetPremiumPlan.GadgetModel.Brand.StringVal,
		sub.GadgetPremiumPlan.GadgetModel.ModelName.StringVal,
	)
}
