package ipaymu_go_api

type Response struct {
	Status int64

	// v2 payment
	Message string        `json:"Message,omitempty"`
	Data    *ResponseData `json:"Data,omitempty"`
}

type ResponseData struct {
	SessionId     string  `json:"SessionId,omitempty"`
	TransactionId int64   `json:"TransactionId,omitempty"`
	ReferenceId   string  `json:"ReferenceId,omitempty"`
	Via           string  `json:"Via,omitempty"`
	Channel       string  `json:"Channel,omitempty"`
	PaymentNo     string  `json:"PaymentNo,omitempty"`
	PaymentName   string  `json:"PaymentName,omitempty"`
	Total         float64 `json:"Total,omitempty"`
	Fee           int64   `json:"Fee,omitempty"`
	Expired       string  `json:"Expired,omitempty"`
	Note          string  `json:"Note,omitempty"`
	Url           string  `json:"Url,omitempty"`
}

type ResponseCheck struct {
	Status int `json:"Status"`
	Data   struct {
		TransactionId  int         `json:"TransactionId"`
		SessionId      string      `json:"SessionId"`
		ReferenceId    interface{} `json:"ReferenceId"`
		RelatedId      int         `json:"RelatedId"`
		Sender         string      `json:"Sender"`
		Receiver       string      `json:"Receiver"`
		Amount         int         `json:"Amount"`
		Fee            int         `json:"Fee"`
		Status         int         `json:"Status"`
		StatusDesc     string      `json:"StatusDesc"`
		Type           int         `json:"Type"`
		TypeDesc       string      `json:"TypeDesc"`
		Notes          string      `json:"Notes"`
		CreatedDate    string      `json:"CreatedDate"`
		ExpiredDate    string      `json:"ExpiredDate"`
		SuccessDate    string      `json:"SuccessDate"`
		SettlementDate string      `json:"SettlementDate"`
	} `json:"Data"`
	Message string `json:"Message"`
}

type ResponseBalance struct {
	Status int `json:"Status"`
	Data   struct {
		Va              string  `json:"Va"`
		MerchantBalance float64 `json:"MerchantBalance"`
		MemberBalance   float64 `json:"MemberBalance"`
	} `json:"Data"`
	Message string `json:"Message"`
}

type ResponseTransaction struct {
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
	Message string `json:"Message"`
	Data    struct {
		Transaction []Transaction `json:"Transaction"`
		Pagination  struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		} `json:"Pagination"`
	} `json:"Data"`
}

type Transaction struct {
	TransactionId  int         `json:"TransactionId"`
	SessionId      *string     `json:"SessionId"`
	ReferenceId    *string     `json:"ReferenceId"`
	RelatedId      int         `json:"RelatedId"`
	Sender         string      `json:"Sender"`
	Receiver       string      `json:"Receiver"`
	Amount         int         `json:"Amount"`
	Fee            int         `json:"Fee"`
	Status         int         `json:"Status"`
	StatusDesc     string      `json:"StatusDesc"`
	PaidStatus     string      `json:"PaidStatus"`
	Type           int         `json:"Type"`
	TypeDesc       string      `json:"TypeDesc"`
	Notes          *string     `json:"Notes"`
	IsEscrow       bool        `json:"IsEscrow"`
	CreatedDate    string      `json:"CreatedDate"`
	ExpiredDate    string      `json:"ExpiredDate"`
	SuccessDate    interface{} `json:"SuccessDate"`
	SettlementDate interface{} `json:"SettlementDate"`
	PaymentChannel string      `json:"PaymentChannel"`
	PaymentCode    string      `json:"PaymentCode"`
	BuyerName      string      `json:"BuyerName"`
	BuyerPhone     string      `json:"BuyerPhone"`
	BuyerEmail     string      `json:"BuyerEmail"`
}

type ResponseListPayment struct {
	Status int `json:"Status"`
	Data   []struct {
		Code          string                 `json:"Code"`
		Description   string                 `json:"Description"`
		Channels      []PaymentChannelDetail `json:"Channels,omitempty"`
		PaymentMethod []PaymentMethodDetail  `json:"PaymentMethod,omitempty"`
	} `json:"Data"`
	Message string `json:"Message"`
}

type PaymentChannelDetail struct {
	Code                 string `json:"Code"`
	Description          string `json:"Description"`
	PaymentIntrucionsDoc string `json:"PaymentIntrucionsDoc"`
	TransactionFee       struct {
		ActualFee     float64 `json:"ActualFee"`
		ActualFeeType string  `json:"ActualFeeType"`
		AdditionalFee float64 `json:"AdditionalFee"`
	} `json:"TransactionFee"`
}

type PaymentMethodDetail struct {
	Code                 string  `json:"Code"`
	Description          string  `json:"Description"`
	PaymentIntrucionsDoc *string `json:"PaymentIntrucionsDoc"`
	TransactionFee       struct {
		ActualFee     float64 `json:"ActualFee"`
		ActualFeeType string  `json:"ActualFeeType"`
		AdditionalFee float64 `json:"AdditionalFee"`
	} `json:"TransactionFee"`
}
