package app

import (
	"fmt"
	"github.com/Thiamath/y2dwdatt/pkg/external/item_service"
	"math"
	"strconv"
	"strings"
)

const ImportedString = "imported"

// TaxService represents the Tax Service class.
type TaxService struct {
	ItemService *item_service.Service

	BasicTax    float64
	ImportTax   float64
	ExemptItems []item_service.ItemLabel
}

// NewTaxService creates a tax service instance with default tax values.
func NewTaxService(itemService *item_service.Service) *TaxService {
	return &TaxService{
		ItemService: itemService,
		BasicTax:    .10,
		ImportTax:   .05,
		ExemptItems: []item_service.ItemLabel{
			item_service.Book,
			item_service.Food,
			item_service.Meds,
		},
	}
}

// Process an input and returns the processed output with taxes and total.
func (t *TaxService) Process(input string) (string, error) {
	lines := strings.Split(input, "\n")
	outputLines := make([]string, 0, len(lines)+2)
	taxes := .0
	total := .0
	for _, line := range lines {
		words := strings.Fields(line)
		num, err := strconv.Atoi(words[0])
		if err != nil {
			return "", fmt.Errorf("error parsing item number on line \"%s\":%w", line, err)
		}
		price, err := strconv.ParseFloat(words[len(words)-1], 64)
		if err != nil {
			return "", fmt.Errorf("error parsing item price on line \"%s\":%w", line, err)
		}
		// remove the number of items and the "at <price>" tail
		words = words[1 : len(words)-2]
		imported, words := t.checkAndCleanImported(words)

		itemName := strings.Join(words, " ")
		item, err := t.ItemService.Get(itemName)
		if err != nil {
			return "", fmt.Errorf("error getting the item \"%s\": %w", itemName, err)
		}
		linePrice := float64(num) * price

		exempted := t.checkExempted(item)
		taxRate := t.BasicTax
		if exempted {
			taxRate = .0
		}
		if imported {
			taxRate += t.ImportTax
		}
		lineTax := linePrice * taxRate
		if lineTax > 0 {
			lineTax = t.roundUpTo005(lineTax)
		}

		linePrice += lineTax

		total += linePrice
		taxes += lineTax
		if imported {
			outputLines = append(outputLines, fmt.Sprintf("%d %s %s: %.2f", num, ImportedString, strings.Join(words, " "), linePrice))
		} else {
			outputLines = append(outputLines, fmt.Sprintf("%d %s: %.2f", num, strings.Join(words, " "), linePrice))
		}
	}
	outputLines = append(outputLines, fmt.Sprintf("Sales Taxes: %.2f", taxes))
	outputLines = append(outputLines, fmt.Sprintf("Total: %.2f", total))
	return strings.Join(outputLines, "\n"), nil
}

func (t *TaxService) roundUpTo005(lineTax float64) float64 {
	fmod := math.Mod(lineTax, .05)
	if fmod > 0 {
		lineTax += .05 - fmod
	}
	return lineTax
}

func (t *TaxService) checkAndCleanImported(words []string) (bool, []string) {
	imported := false
	outputWords := make([]string, 0)
	for _, word := range words {
		if word == ImportedString {
			imported = true
		} else {
			outputWords = append(outputWords, word)
		}
	}
	return imported, outputWords
}

func (t *TaxService) checkExempted(item *item_service.Item) bool {
	exempted := false
	for _, exemptItem := range t.ExemptItems {
		if item.Label == exemptItem {
			exempted = true
			break
		}
	}
	return exempted
}
