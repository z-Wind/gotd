package gotd

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// NewTransactionHistoryService https://developer.tdameritrade.com/transaction-history/apis
// APIs to access transaction history on the account
func NewTransactionHistoryService(s *Service) *TransactionHistoryService {
	rs := &TransactionHistoryService{s: s}
	return rs
}

// TransactionHistoryService https://developer.tdameritrade.com/transaction-history/apis
// APIs to access transaction history on the account
type TransactionHistoryService struct {
	s *Service
}

// GetTransaction https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions/%7BtransactionId%7D-0
// Transaction for a specific account.
// 感覺 api 有問題，給正確的 transactionID 卻找不到
func (r *TransactionHistoryService) GetTransaction(accountID, transactionID string) *TransactionHistoryGetTransactionCall {
	c := &TransactionHistoryGetTransactionCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:     accountID,
		transactionID: transactionID,
	}
	return c
}

// TransactionHistoryGetTransactionCall https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions/%7BtransactionId%7D-0
// Transaction for a specific account.
// 感覺 api 有問題，給正確的 transactionID 卻找不到
type TransactionHistoryGetTransactionCall struct {
	DefaultCall

	accountID, transactionID string
}

func (c *TransactionHistoryGetTransactionCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "transactions", c.transactionID)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TransactionHistoryGetTransactionCall) Do() (*Transaction, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &Transaction{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	return ret, nil

}

// GetTransactionList https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions-0
// Transactions for a specific account.
func (r *TransactionHistoryService) GetTransactionList(accountID string) *TransactionHistoryGetTransactionListCall {
	c := &TransactionHistoryGetTransactionListCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}

	return c
}

// TransactionHistoryGetTransactionListCall https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions-0
// Transactions for a specific account.
type TransactionHistoryGetTransactionListCall struct {
	DefaultCall

	accountID string
}

// Type https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions-0
// Only transactions with the specified type will be returned.
func (c *TransactionHistoryGetTransactionListCall) Type(t string) *TransactionHistoryGetTransactionListCall {
	c.urlParams.Set("type", t)
	return c
}

// Symbol https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions-0
// Only transactions with the specified symbol will be returned.
func (c *TransactionHistoryGetTransactionListCall) Symbol(symbol string) *TransactionHistoryGetTransactionListCall {
	c.urlParams.Set("symbol", symbol)
	return c
}

// StartDate https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions-0
// Only transactions after the Start Date will be returned.
// Note: The maximum date range is one year. Valid ISO-8601 formats are :
// yyyy-MM-dd.
func (c *TransactionHistoryGetTransactionListCall) StartDate(date time.Time) *TransactionHistoryGetTransactionListCall {
	c.urlParams.Set("startDate", date.Format("2006-01-02"))
	return c
}

// EndDate https://developer.tdameritrade.com/transaction-history/apis/get/accounts/%7BaccountId%7D/transactions-0
// Only transactions before the End Date will be returned.
// Note: The maximum date range is one year. Valid ISO-8601 formats are :
// yyyy-MM-dd.
func (c *TransactionHistoryGetTransactionListCall) EndDate(date time.Time) *TransactionHistoryGetTransactionListCall {
	c.urlParams.Set("endDate", date.Format("2006-01-02"))
	return c
}

func (c *TransactionHistoryGetTransactionListCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "transactions")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *TransactionHistoryGetTransactionListCall) Do() (*TransactionList, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &TransactionList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*Transaction)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.Transactions = *target
	return ret, nil

}
