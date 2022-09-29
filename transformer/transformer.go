package transformer

import (
	"math/rand"
	"se_cli/model"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func TransformDataToSeProducts(source model.Response) []model.Product {
	products := make([]model.Product, 0)


	for _, data := range source.Result {
		variants := make([]model.Variant, 0)

		for _, option := range data.OptionColor {
			variants = append(variants, model.Variant{
				VariantName: option.DisplayName,
			})
		}


		products = append(products, model.Product{
			Name: data.Name,
			Code: RandStringRunes(10),
			Description: data.ShortDescription,
			Variants: variants,
		})
	}


	return products
}