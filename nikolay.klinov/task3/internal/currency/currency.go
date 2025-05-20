package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ExchangeRate struct {
	NumberCode int     `json:"numberCode"`
	Symbol     string  `json:"symbol"`
	Rate       float64 `json:"rate"`
}

type CurrencyList struct {
	Items []CurrencyItem `xml:"Valute"`
}

type CurrencyItem struct {
	NumberCode string `xml:"NumCode"`
	Symbol     string `xml:"CharCode"`
	RateValue  string `xml:"Value"`
}

func ParseRatesFromXML(sourcePath string) ([]ExchangeRate, error) {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewReader(data)
	xmlParser := xml.NewDecoder(buffer)
	xmlParser.CharsetReader = charset.NewReaderLabel

	var xmlData CurrencyList
	err = xmlParser.Decode(&xmlData)
	if err != nil {
		return nil, err
	}

	var rates []ExchangeRate
	for _, item := range xmlData.Items {
		cleanRate := strings.Replace(item.RateValue, ",", ".", -1)
		parsedRate, err := strconv.ParseFloat(cleanRate, 64)
		if err != nil {
			return nil, err
		}

		code, err := strconv.Atoi(item.NumberCode)
		if err != nil {
			return nil, err
		}

		rates = append(rates, ExchangeRate{
			NumberCode: code,
			Symbol:     item.Symbol,
			Rate:       parsedRate,
		})
	}
	return rates, nil
}

func SortByRate(rates []ExchangeRate) []ExchangeRate {
	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Rate > rates[j].Rate
	})
	return rates
}

func ExportToJSON(rates []ExchangeRate, destinationPath string) error {
	outputFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	jsonWriter := json.NewEncoder(outputFile)
	jsonWriter.SetIndent("", "    ")
	return jsonWriter.Encode(rates)
}
