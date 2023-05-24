package database

type CStorage interface {
	Get(key string) (value any, err error)
	GetOne(T any) (value, err error)
	Insert(data []byte, key string) error
	Update(T any) error
	Delete(T any) error
}

type DBStorage interface {
	GetUsers() (users []*User, err error)
	GetOneUserById(id int64) (value any, err error)
	InsertUser(user *User) error
	UpdateUser() error
	DeleteUser() error
	GetStocksFromUser(id int64) (value []Stock, err error)
	GetStockAndCheckLastUpdate(symbol string) (value bool, err error)
	UpdateStockInfo(symbol string, stock *Stock) error
	GetStockInfo(symbol string) (value *Stock, err error)
	InsertNewStockInfo(stock *Stock) error
	InsertNewLinkRss(link string, userId int64) error
	GetAllLinkRss() (value []string, err error)
	GetAllLinkRssForUser(userId int64) (value []string, err error)
	DeleteLinkRss(link string) error
	DeleteStockFromUserById() error
	DeleteStock() error
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Stocks   any    `json:"stocks"`
}

type Stock struct {
	ID                         int64  `json:"id"`
	Symbol                     string `json:"Symbol"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	CIK                        string `json:"CIK"`
	Exchange                   string `json:"Exchange"`
	Currency                   string `json:"Currency"`
	Country                    string `json:"Country"`
	Sector                     string `json:"Sector"`
	Industry                   string `json:"Industry"`
	Address                    string `json:"Address"`
	FiscalYearEnd              string `json:"FiscalYearEnd"`
	LatestQuarter              string `json:"LatestQuarter"`
	MarketCapitalization       string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookValue                  string `json:"BookValue"`
	DividendPerShare           string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM"`
	RevenueTTM                 string `json:"RevenueTTM"`
	GrossProfitTTM             string `json:"GrossProfitTTM"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice"`
	TrailingPE                 string `json:"TrailingPE"`
	ForwardPE                  string `json:"ForwardPE"`
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	WeekHigh                   string `json:"WeekHigh"`
	WeekLow                    string `json:"WeekLow"`
	DayMovingAverage           string `json:"DayMovingAverage50"`
	DayMovingAverage200        string `json:"DayMovingAverage200"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	SharesFloat                string `json:"SharesFloat"`
	SharesShort                string `json:"SharesShort"`
	SharesShortPriorMonth      string `json:"SharesShortPriorMonth"`
	ShortRatio                 string `json:"ShortRatio"`
	ShortPercentOutstanding    string `json:"ShortPercentOutstanding"`
	ShortPercentFloat          string `json:"ShortPercentFloat"`
	PercentInsiders            string `json:"PercentInsiders"`
	PercentInstitutions        string `json:"PercentInstitutions"`
	ForwardAnnualDividendRate  string `json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield string `json:"ForwardAnnualDividendYield"`
	PayoutRatio                string `json:"PayoutRatio"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
	LastSplitFactor            string `json:"LastSplitFactor"`
	LastSplitDate              string `json:"LastSplitDate"`
	UpdatedAt                  string `json:"UpdatedAt"`
}
