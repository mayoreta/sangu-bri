package bri

type TokenResponse struct {
	AccessToken string   `json:"access_token"`
	ExpiredTime string   `json:"expires_in"`
	ProductList []string `json:"api_product_list_json"`
}

type VaResponse struct {
	Status              bool   `json:"status"`
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
	ErrDesc             string `json:"errDesc"`
	Data                VaData `json:"data"`
}

type VaData struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Name            string `json:"nama"`
	Amount          string `json:"amount"`
	Description     string `json:"keterangan"`
	StatusBayar     string `json:"statusBayar"`
	ExpiredDate     string `json:"expiredDate"`
}

type VaReportResponse struct {
	Status       bool           `json:"status"`
	ResponseCode string         `json:"responseCode"`
	Description  string         `json:"responseDescription"`
	ErrDesc      string         `json:"errDesc"`
	Data         []VaReportData `json:"data"`
}

type VaReportData struct {
	BrivaNo     string `json:"brivaNo"`
	CustCode    string `json:"custCode"`
	Nama        string `json:"nama"`
	Amount      string `json:"amount"`
	Description string `json:"keterangan"`
	PaymentDate string `json:"paymentDate"`
	TellerId    string `json:"tellerid"`
	AccountNo   string `json:"no_rek"`
}

// CardTokenOTPResponse defines response for direct debit - create card token OTP
type CardTokenOTPResponse struct {
	Body CardTokenOTPResponseData `json:"body"`
	ErrorResponse
}

// CardTokenOTPResponseData defines data response for direct debit - create card token OTP
type CardTokenOTPResponseData struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

// CardTokenOTPVerifyResponse defines response for direct debit - create card token OTP verify
type CardTokenOTPVerifyResponse struct {
	Body CardTokenOTPVerifyResponseData `json:"body"`
	ErrorResponse
}

// CardTokenOTPVerifyResponseData defines data response for direct debit - create card token OTP verify
type CardTokenOTPVerifyResponseData struct {
	Status           string                 `json:"status"`
	PhoneNumber      string                 `json:"phone_number"`
	DeviceID         string                 `json:"device_id"`
	CardToken        string                 `json:"card_token"`
	Location         Location               `json:"location"`
	Last4            string                 `json:"last4"`
	Email            string                 `json:"email"`
	CardType         string                 `json:"card_type"`
	LimitTransaction string                 `json:"limit_transaction"`
	Metadata         map[string]interface{} `json:"metadata"`
}

// ErrorResponse defines response data if error request.Example:
//
//	{
//	    "error": {
//	        "code": "0920",
//	        "message": "Expired OTP"
//	    },
//	    "status_code": 400,
//	    "status": {
//	        "code": "0602",
//	        "desc": "Invalid signature"
//	    },
//	    "recorded_at": "2020-03-12T03:58:46Z"
//	}
type ErrorResponse struct {
	Error      ErrorDetail `json:"error"`
	StatusCode int         `json:"status_code"`
	Status     ErrorStatus `json:"status"`
}

// ErrorDetail defines response error detail. Example:
//
//	{
//	    "error": {
//	        "code": "0920",
//	        "message": "Expired OTP"
//	    }
//	}
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorStatus defines error data if unauthorized. Example:
//
//	{
//	    "status": {
//	        "code": "0602",
//	        "desc": "Invalid signature"
//	    }
//	}
type ErrorStatus struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

// PaymentChargeResponse defines response for direct debit - create payment charge [using OTP or not]
type PaymentChargeResponse struct {
	Body PaymentChargeResponseData `json:"body"`
	ErrorResponse
}

// PaymentChargeResponseData defines data response for direct debit - create payment charge [using OTP or not]
type PaymentChargeResponseData struct {
	Status        string                 `json:"status"`
	ChargeToken   string                 `json:"charge_token"`
	PaymentID     string                 `json:"payment_id"`
	Amount        string                 `json:"amount"`
	Currency      string                 `json:"currency"`
	Remarks       string                 `json:"remarks"`
	DeviceID      string                 `json:"device_id"`
	PaymentStatus string                 `json:"payment_status"`
	Location      Location               `json:"location"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// DeleteCardTokenResponse defines response for direct debit - delete card token
type DeleteCardTokenResponse struct {
	Body DeleteCardTokenResponseData `json:"body"`
	ErrorResponse
}

// DeleteCardTokenResponseData defines data response for direct debit - delete card token
type DeleteCardTokenResponseData struct {
	Status string `json:"status"`
}

// ChargeDetailResponse defines response for direct debit - charge detail
type ChargeDetailResponse struct {
	Body ChargeDetailResponseData `json:"body"`
	ErrorResponse
}

// ChargeDetailResponseData defines data response for direct debit - charge detail
type ChargeDetailResponseData struct {
	Status          string                 `json:"status"`
	Amount          string                 `json:"amount"`
	Currency        string                 `json:"currency"`
	PaymentID       string                 `json:"payment_id"`
	RemarksMerchant string                 `json:"remarks_merchant"`
	PaymentStatus   string                 `json:"payment_status"`
	RefundHistory   []RefundResponseData   `json:"refund_history"`
	DeviceID        string                 `json:"device_id"`
	Location        Location               `json:"location"`
	Metadata        map[string]interface{} `json:"metadata"`
}

// RefundResponseData defines data response for direct debit - refund
type RefundResponseData struct {
	Status       string                 `json:"status"`
	RefundID     string                 `json:"refund_id"`
	PaymentID    string                 `json:"payment_id"`
	Amount       string                 `json:"amount"`
	Fee          string                 `json:"fee"`
	Currency     string                 `json:"currency"`
	Reason       string                 `json:"reason"`
	RefundStatus string                 `json:"refund_status"`
	DeviceID     string                 `json:"device_id"`
	Location     Location               `json:"location"`
	Metadata     map[string]interface{} `json:"metadata"`
	Date         string                 `json:"date"`
}

// RefundResponse defines response for direct debit - refund
type RefundResponse struct {
	Body RefundResponseData `json:"body"`
	ErrorResponse
}

// MutationRespCode is BRI API response code standard ( more at https://developers.bri.co.id/id/docs/account-statementv20 )
const MutationRespCodeSuccess = "0000"

type MutationResponse struct {
	ResponseCode        string         `json:"responseCode"`
	ResponseDescription string         `json:"responseDescription"`
	ErrDesc             string         `json:"errDesc"`
	Data                []MutationData `json:"data"`
}

type MutationData struct {
	TransactionTime string `json:"transactionTime"`
	DebitAmount     string `json:"debitAmount"`
	CreditAmount    string `json:"creditAmount"`
	TypeAmount      string `json:"typeAmount"`
	Remark          string `json:"remark"`
	StartBalance    string `json:"startBalance"`
	EndBalance      string `json:"endBalance"`
}
