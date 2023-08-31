package constants

const (
	ApiHostMainnet = "https://api.dydx.exchange"
	ApiHostRopsten = "https://api.stage.dydx.exchange"
	WsHostMainnet  = "wss://api.dydx.exchange/v3/ws"
	WsHostRopsten  = "wss://api.stage.dydx.exchange/v3/ws"
)

const (
	SignatureTypeNoPrepend   = 0
	SignatureTypeDecimal     = 1
	SignatureTypeHexadecimal = 2
)

const (
	Domain                       = "dYdX"
	Version                      = "1.0"
	Eip712DomainStringNoContract = "EIP712Domain(string name,string version,uint256 chainId)"
)

const Eip712OnboardingActionStructString = "dYdX(string action,string onlySignOn)"

const Eip712OnboardingActionStructStringTestnet = "dYdX(string action)"
const Eip712StructName = "dYdX"
const OnlySignOnDomainMainnet = "https://trade.dydx.exchange"

const (
	OffChainOnboardingAction    = "dYdX Onboarding"
	OffChainKeyDerivationAction = "dYdX STARK Key"
)

const (
	NetworkIdMainnet = 1
	NetworkIdGoerli  = 5
)

const (
	OPEN        = "OPEN"
	CLOSED      = "CLOSED"
	LIQUIDATED  = "LIQUIDATED"
	LIQUIDATION = "LIQUIDATION"
	UNTRIGGERED = "UNTRIGGERED"
)

const (
	MARKET       = "MARKET"
	LIMIT        = "LIMIT"
	STOP         = "STOP"
	STOPLIMIT    = "STOP_LIMIT"
	TRAILINGSTOP = "TRAILING_STOP"
	TAKEPROFIT   = "TAKE_PROFIT"
)

const (
	BUY  = "BUY"
	SELL = "SELL"
)

const (
	LONG  = "LONG"
	SHORT = "SHORT"
)

const (
	TimeInForceGtt = "GTT"
	TimeInForceFok = "FOK"
	TimeInForceIoc = "IOC"
)

const (
	OrderStatusPending     = "PENDING"
	OrderStatusOpen        = "OPEN"
	OrderStatusFilled      = "FILLED"
	OrderStatusCanceled    = "CANCELED"
	OrderStatusUntriggered = "UNTRIGGERED"
)

const (
	Resolution1D     = "1DAY"
	Resolution4HOURS = "4HOURS"
	Resolution1HOUR  = "1HOUR"
	Resolution30MINS = "30MINS"
	Resolution15MINS = "15MINS"
	Resolution5MINS  = "5MINS"
	Resolution1MIN   = "1MIN"
)
