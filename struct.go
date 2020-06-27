package main

import (
	"net/http"
)

type PhemexClient struct {
	Client *http.Client
	Api    string
	Secret []byte
}

type Market struct {
	Symbol             string `json:"symbol"`
	UnderlyingSymbol   string `json:"underlyingSymbol"`
	QuoteCurrency      string `json:"quoteCurrency"`
	SettlementCurrency string `json:"settlementCurrency"`
	MaxOrderQty        int64  `json:"maxOrderQty"`
	MaxPriceEp         int64  `json:"maxPriceEp"`
	LotSize            int64  `json:"lotSize"`
	TickSize           int64  `json:"tickSize"`
	ContractSize       string `json:"contractSize"`
	PriceScale         int64  `json:"priceScale"`
	RatioScale         int64  `json:"ratioScale"`
	ValueScale         int64  `json:"valueScale"`
	DefaultLeverage    int64  `json:"defaultLeverage"`
	MaxLeverage        int64  `json:"maxLeverage"`
	InitMarginEr       int64  `json:"initMarginEr"`
	MaintMarginEr      int64  `json:"maintMarginEr"`
	DefaultRiskLimitEv int64  `json:"defaultRiskLimitEv"`
	Deleverage         bool   `json:"deleverage"`
	MakerFeeRateEr     int64  `json:"makerFeeRateEr"`
	TakerFeeRateEr     int64  `json:"takerFeeRateEr"`
	FundingInterval    int64  `json:"fundingInterval"`
	Description        string `json:"description"`
	Type               string `json:"type"`
}
