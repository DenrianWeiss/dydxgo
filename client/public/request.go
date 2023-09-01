package public

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (p *Public) GetMarkets(marketID string) (*MarketsResponse, error) {
	u := url.Values{}
	if marketID != "" {
		u.Add("market", marketID)
	}
	res, err := p.get("markets", u)
	if err != nil {
		return nil, err
	}
	m := &MarketsResponse{}
	if err := json.Unmarshal(res, m); err != nil {
		return nil, err
	}

	return m, nil
}

func (p *Public) GetTrades(param *TradesParam) (*TradesResponse, error) {
	u, err := query.Values(param)
	if err != nil {
		return nil, errors.New("error when changed struct to query")
	}

	res, err := p.get(fmt.Sprintf("trades/%s", param.MarketID), u)
	if err != nil {
		return nil, err
	}
	t := &TradesResponse{}
	if err := json.Unmarshal(res, t); err != nil {
		return nil, err
	}

	return t, nil
}

// GetCandles response head of the array is currently
func (p *Public) GetCandles(param *CandlesParam) (*CandlesResponse, error) {
	u, err := query.Values(param)
	if err != nil {
		return nil, errors.New("error when changed struct to query")
	}

	res, err := p.get(fmt.Sprintf("candles/%s", param.Market), u)
	if err != nil {
		return nil, err
	}
	t := &CandlesResponse{}
	if err := json.Unmarshal(res, t); err != nil {
		return nil, err
	}

	return t, nil
}

func (p *Public) GetHistoricalFunding(param *HistoricalFundingsParam) (*HistoricalFundingsResponse, error) {
	u, err := query.Values(param)
	if err != nil {
		return nil, errors.New("error when changed struct to query")
	}

	res, err := p.get(fmt.Sprintf("historical-funding/%s", param.Market), u)
	if err != nil {
		return nil, err
	}
	t := &HistoricalFundingsResponse{}
	if err := json.Unmarshal(res, t); err != nil {
		return nil, err
	}

	return t, nil
}
