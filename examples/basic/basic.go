package main

import (
	"strings"
	"time"

	"github.com/ajef/ibapi"
)

const (
	host = "localhost"
	port = 7497
)

var orderID int64

func nextID() int64 {
	orderID++
	return orderID
}

func main() {
	// We set logger for pretty logs to console
	log := ibapi.Logger()
	//ibapi.SetLogLevel(int(zerolog.DebugLevel))
	ibapi.SetConsoleWriter()
	// ibapi.SetConnectionTimeout(1 * time.Second)

	// IB CLient
	ib := ibapi.NewEClient(nil)

	if err := ib.Connect(host, port, 5); err != nil { //rand.Int63n(999999)
		log.Error().Err(err).Msg("Connect")
		return
	}

	// Add a short delay to allow the connection to stabilize
	time.Sleep(100 * time.Millisecond)
	log.Info().Msg("Waited for connection to stabilize")

	// ib.SetConnectionOptions("+PACEAPI")

	// Logger test
	// log.Trace().Interface("Log level", log.GetLevel()).Msg("Logger Trace")
	// log.Debug().Interface("Log level", log.GetLevel()).Msg("Logger Debug")
	// log.Info().Interface("Log level", log.GetLevel()).Msg("Logger Info")
	// log.Warn().Interface("Log level", log.GetLevel()).Msg("Logger Warn")
	// log.Error().Interface("Log level", log.GetLevel()).Msg("Logger Error")

	time.Sleep(1 * time.Second)
	log.Info().Bool("Is connected", ib.IsConnected()).Int("Server Version", ib.ServerVersion()).Str("TWS Connection time", ib.TWSConnectionTime()).Msg("Connection details")

	time.Sleep(1 * time.Second)
	ib.ReqCurrentTime()

	// ########## account ##########
	// ib.ReqManagedAccts()

	// ib.ReqAutoOpenOrders(false) // Only from clientID = 0
	// ib.ReqAutoOpenOrders(false)
	// ib.ReqAccountUpdates(true, "")
	// ib.ReqAllOpenOrders()
	// ib.ReqPositions()
	// ib.ReqCompletedOrders(false)
	// ib.ReqExecutions(690, ibapi.NewExecutionFilter())

	tags := []string{"AccountType", "NetLiquidation", "TotalCashValue", "SettledCash",
		"sAccruedCash", "BuyingPower", "EquityWithLoanValue",
		"PreviousEquityWithLoanValue", "GrossPositionValue", "ReqTEquity",
		"ReqTMargin", "SMA", "InitMarginReq", "MaintMarginReq", "AvailableFunds",
		"ExcessLiquidity", "Cushion", "FullInitMarginReq", "FullMaintMarginReq",
		"FullAvailableFunds", "FullExcessLiquidity", "LookAheadNextChange",
		"LookAheadInitMarginReq", "LookAheadMaintMarginReq",
		"LookAheadAvailableFunds", "LookAheadExcessLiquidity",
		"HighestSeverity", "DayTradesRemaining", "Leverage", "$LEDGER:ALL"}
	id := nextID()
	ib.ReqAccountSummary(id, "All", strings.Join(tags, ","))
	time.Sleep(10 * time.Second)
	ib.CancelAccountSummary(id)

	// time.Sleep(1 * time.Second)
	// ib.ReqFamilyCodes()
	// time.Sleep(1 * time.Second)
	// ib.ReqScannerParameters()
	// ########## market data ##########
	//eurusd := &ibapi.Contract{Symbol: "EUR", SecType: "CASH", Currency: "USD", Exchange: "IDEALPRO"}
	//id := nextID()
	// ib.ReqMktData(id, eurusd, "", false, false, nil)
	// time.Sleep(4 * time.Second)
	// ib.CancelMktData(id)

	// ########## real time bars ##########
	// aapl := &ibapi.Contract{ConID: 265598, Symbol: "AAPL", SecType: "STK", Exchange: "SMART"}
	// id := nextID()
	// ib.ReqRealTimeBars(id, aapl, 5, "TRADES", false, nil)
	// time.Sleep(10 * time.Second)
	// ib.CancelRealTimeBars(id)

	//  ########## contract ##########
	// ib.ReqContractDetails(nextID(), aapl)
	// ib.ReqMatchingSymbols(nextID(), "ibm")

	// ########## orders ##########
	// id := nextID()
	// eurusd := &ibapi.Contract{Symbol: "EUR", SecType: "CASH", Currency: "USD", Exchange: "IDEALPRO"}
	// limitOrder := ibapi.LimitOrder("BUY", ibapi.StringToDecimal("20000"), 1.08)
	// ib.PlaceOrder(id+102, eurusd, limitOrder)

	// time.Sleep(4 * time.Second)
	// ib.CancelOrder(id, ibapi.NewOrderCancel())
	// time.Sleep(4 * time.Second)
	// ib.ReqGlobalCancel(ibapi.NewOrderCancel())
	// Real time bars
	// duration := "60 S"
	// barSize := "5 secs"
	// whatToShow := "MIDPOINT" // "TRADES", "MIDPOINT", "BID" or "ASK"
	// ib.ReqHistoricalData(id, eurusd, "", duration, barSize, whatToShow, true, 1, true, nil)

	// ########## orders ##########
	// id := nextID()
	// opt := ibapi.NewContract()
	// opt.Symbol = "GOOG"
	// opt.SecType = "OPT"
	// opt.Exchange = "SMART"
	// opt.Currency = "USD"
	// opt.LastTradeDateOrContractMonth = "202512"
	// opt.Strike = 150
	// opt.Right = "C"
	// opt.Multiplier = "100"
	// ib.CalculateOptionPrice(id, opt, 0.25, 150, nil)
	// time.Sleep(3 * time.Second)
	// ib.CalculateImpliedVolatility(nextID(), opt, 14.35, 150, nil)
	// time.Sleep(10 * time.Second)
	//ib.CancelHistoricalData(id)

	time.Sleep(5 * time.Second)
	err := ib.Disconnect()
	if err != nil {
		log.Error().Err(err).Msg("Disconnect")
	}
	log.Info().Msg("Bye!!!!")
}
