package controllers

import (
	"errors"
	"strconv"

	"github.com/DaveSaah/risky/model"
	"github.com/labstack/echo/v4"
)

// Validate validates user input from the form
func Validate(c echo.Context) (model.TradeInfo, error) {
	var tradeInfo model.TradeInfo

	cap, err := strconv.ParseFloat(c.FormValue("capital"), 64)
	if err != nil || cap < 1 {
		return tradeInfo, errors.New(
			"invalid value for Capital. Enter a number greater than 1",
		)
	}

	lev, err := strconv.Atoi(c.FormValue("leverage"))
	if err != nil || lev < 1 {
		return tradeInfo, errors.New(
			"invalid value for Leverage. Enter a number greater than 1",
		)
	}

	risk, err := strconv.ParseFloat(c.FormValue("risk"), 64)
	if err != nil || !(risk > 0 && risk < 100) {
		return tradeInfo, errors.New(
			"invalid value for Risk. Either enter a number or pick a value between 1 and 100",
		)
	}

	trade_count, err := strconv.Atoi(c.FormValue("trade_count"))
	if err != nil || trade_count < 1 {
		return tradeInfo, errors.New(
			"invalid value for Number of trades. Enter a number greater than 1",
		)
	}

	rr, err := strconv.ParseFloat(c.FormValue("rr"), 64)
	if err != nil || rr < 1 {
		return tradeInfo, errors.New(
			"invalid value for Risk/Reward ratio. Enter a number greater than 1",
		)
	}

	asset, err := strconv.Atoi(c.FormValue("asset"))
	if err != nil || (asset < 0 || asset >= 4) {
		return tradeInfo, errors.New(
			"invalid value for asset. Your chosen pair is not supported",
		)
	}

	tradeInfo = model.TradeInfo{
		Capital:     cap,
		Leverage:    lev,
		RiskPercent: risk,
		TradeCount:  trade_count,
		RR:          rr,
		Asset:       model.Pair(asset),
	}

	return tradeInfo, nil
}
