package private

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/denrianweiss/dydxgo/utils"
	"github.com/yanue/starkex"
	"net/url"
	"strconv"
)

func (p *Private) GetUsers() (*UsersResponse, error) {
	res, err := p.get("users", nil)
	if err != nil {
		return nil, err
	}

	result := &UsersResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *Private) GetAccount(ethereumAddress string) (*AccountResponse, error) {
	if ethereumAddress == "" {
		ethereumAddress = p.Address.String()
	}
	uri := fmt.Sprintf("accounts/%s", utils.GetAccountId(ethereumAddress))
	res, err := p.get(uri, nil)
	if err != nil {
		return nil, err
	}

	result := &AccountResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *Private) CreateOrder(input *ApiOrder, positionId int64) (*OrderResponse, error) {
	orderSignParam := starkex.OrderSignParam{
		NetworkId:  int(p.NetworkId),
		PositionId: positionId,
		Market:     input.Market,
		Side:       input.Side,
		HumanSize:  input.Size,
		HumanPrice: input.Price,
		LimitFee:   input.LimitFee,
		ClientId:   input.ClientId,
		Expiration: input.Expiration,
	}
	signature, err := starkex.OrderSign(p.StarkPrivateKey, orderSignParam)
	if err != nil {
		return nil, errors.New("sign error")
	}
	input.Signature = signature

	res, err := p.post("orders", input)
	if err != nil {
		return nil, err
	}

	result := &OrderResponse{}
	if err = json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetPositions(market string) (*PositionResponse, error) {
	params := url.Values{}
	if market != "" {
		params.Add("market", market)
	}

	res, err := p.get("positions", params)
	if err != nil {
		return nil, err
	}

	result := &PositionResponse{}
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetOrders(input *OrderQueryParam) (*OrderListResponse, error) {
	data, err := p.get("orders", input.ToParams())
	if err != nil {
		return nil, err
	}

	result := &OrderListResponse{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) CancelOrder(orderId string) (*CancelOrderResponse, error) {
	data, err := p.delete("orders/"+orderId, nil)
	if err != nil {
		return nil, err
	}

	result := &CancelOrderResponse{}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) CancelOrders(market string) (*CancelOrdersResponse, error) {
	values := url.Values{}
	if market != "" {
		values.Add("market", market)
	}

	res, err := p.delete("orders", values)
	if err != nil {
		return nil, err
	}
	result := &CancelOrdersResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetOrderById(orderId string) (*OrderResponse, error) {
	res, err := p.get("orders/"+orderId, nil)
	if err != nil {
		return nil, err
	}

	result := &OrderResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) Withdraw(param *WithdrawalParam) (*WithdrawResponse, error) {
	signature, err := starkex.WithdrawSign(p.StarkPrivateKey, starkex.WithdrawSignParam{
		NetworkId:   int(p.NetworkId),
		ClientId:    param.ClientID,
		PositionId:  int64(utils.ToFloat(param.PositionId)),
		HumanAmount: param.Amount,
		Expiration:  param.Expiration,
	})
	if err != nil {
		return nil, errors.New("sign error")
	}
	param.Signature = signature
	res, err := p.post("withdrawals", param)

	if err != nil {
		return nil, err
	}

	result := &WithdrawResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) WithdrawFast(param *WithdrawalFastParam, lpStarkKey string, positionId string) (*WithdrawResponse, error) {
	convertedId, _ := strconv.Atoi(positionId)
	convertedLpId, _ := strconv.Atoi(param.LpPositionId)
	signature, err := starkex.TransferSign(p.StarkPrivateKey, starkex.TransferSignParam{
		NetworkId:          int(p.NetworkId),
		SenderPositionId:   int64(convertedId),
		ReceiverPositionId: int64(convertedLpId),
		ReceiverPublicKey:  lpStarkKey,
		ReceiverAddress:    param.ToAddress,
		CreditAmount:       param.CreditAmount,
		DebitAmount:        param.DebitAmount,
		Expiration:         param.Expiration,
		ClientId:           param.ClientID,
	})

	if err != nil {
		return nil, errors.New("sign error")
	}
	param.Signature = signature

	res, err := p.post("fast-withdrawals", param)
	if err != nil {
		return nil, err
	}

	result := &WithdrawResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetFills(param *FillsParam) (*FillsResponse, error) {
	u := utils.ToValues(param)
	res, err := p.get("fills", u)
	if err != nil {
		return nil, err
	}

	result := &FillsResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetFundingPayments(param *FundingPaymentsParam) (*FundingPaymentsResponse, error) {
	u := utils.ToValues(param)
	res, err := p.get("funding", u)
	if err != nil {
		return nil, err
	}

	result := &FundingPaymentsResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetHistoricalPnL(param *HistoricalPnLParam) (*HistoricalPnLResponse, error) {
	u := utils.ToValues(param)
	res, err := p.get("historical-pnl", u)
	if err != nil {
		return nil, err
	}

	result := &HistoricalPnLResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetTradingRewards(param *TradingRewardsParam) (*TradingRewardsResponse, error) {
	u := utils.ToValues(param)
	res, err := p.get("rewards/weight", u)
	if err != nil {
		return nil, err
	}

	result := &TradingRewardsResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetTransfers(param *TransferQueryParam) (*TransferQueryResponse, error) {
	u := utils.ToValues(param)
	res, err := p.get("transfers", u)
	if err != nil {
		return nil, err
	}

	result := &TransferQueryResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}

func (p *Private) GetRegistration() (*RegisterResponse, error) {
	res, err := p.get("registration", nil)
	if err != nil {
		return nil, err
	}

	result := &RegisterResponse{}
	if err := json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return result, nil
}
