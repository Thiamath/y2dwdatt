package app_test

import (
	"github.com/Thiamath/y2dwdatt/app"
	"github.com/Thiamath/y2dwdatt/pkg/external/item_service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App package", func() {
	Context("Tax service", func() {
		taxService := app.NewTaxService(ItemServiceMock)
		It(`Should be able to process input:
1 book at 12.49
1 music CD at 14.99
1 chocolate bar at 0.85`, func() {
			// Given input alpha
			// When passing through the taxes service
			output, err := taxService.Process(inputAlpha)
			// Then should not be errors
			Expect(err).To(BeNil())
			// And the output must be alpha output
			Expect(output).To(Equal(outputAlpha))
		})
		It(`Should be able to process input:
1 imported box of chocolates at 10.00
1 imported bottle of perfume at 47.50`, func() {
			// Given input beta
			// When passing through the taxes service
			output, err := taxService.Process(inputBeta)
			// Then should not be errors
			Expect(err).To(BeNil())
			// And the output must be beta output
			Expect(output).To(Equal(outputBeta))
		})
		It(`Should be able to process input:
1 imported bottle of perfume at 27.99
1 bottle of perfume at 18.99
1 packet of headache pills at 9.75
1 box of imported chocolates at 11.25`, func() {
			// Given input gamma
			// When passing through the taxes service
			output, err := taxService.Process(inputGamma)
			// Then should not be errors
			Expect(err).To(BeNil())
			// And the output must be gamma output
			Expect(output).To(Equal(outputGamma))
		})
	})
})

var (
	inputAlpha = `1 book at 12.49
1 music CD at 14.99
1 chocolate bar at 0.85`

	inputBeta = `1 imported box of chocolates at 10.00
1 imported bottle of perfume at 47.50`

	inputGamma = `1 imported bottle of perfume at 27.99
1 bottle of perfume at 18.99
1 packet of headache pills at 9.75
1 box of imported chocolates at 11.25`

	outputAlpha = `1 book: 12.49
1 music CD: 16.49
1 chocolate bar: 0.85
Sales Taxes: 1.50
Total: 29.83`

	outputBeta = `1 imported box of chocolates: 10.50
1 imported bottle of perfume: 54.65
Sales Taxes: 7.65
Total: 65.15`

	outputGamma = `1 imported bottle of perfume: 32.19
1 bottle of perfume: 20.89
1 packet of headache pills: 9.75
1 imported box of chocolates: 11.85
Sales Taxes: 6.70
Total: 74.68`
)

var ItemServiceMock = &item_service.Service{
	FictionalDatabase: map[string]*item_service.Item{
		"book": {
			Name:  "book",
			Label: item_service.Book,
		},
		"music CD": {
			Name: "music CD",
		},
		"chocolate bar": {
			Name:  "chocolate bar",
			Label: item_service.Food,
		},
		"box of chocolates": {
			Name:  "box of chocolates",
			Label: item_service.Food,
		},
		"chocolates": {
			Name:  "chocolates",
			Label: item_service.Food,
		},
		"bottle of perfume": {
			Name: "bottle of perfume",
		},
		"packet of headache pills": {
			Name:  "packet of headache pills",
			Label: item_service.Meds,
		},
	},
}
