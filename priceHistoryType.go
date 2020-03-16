package gotd

import (
	"strconv"
	"time"
)

// Candle https://developer.tdameritrade.com/price-history/apis
type Candle struct {
	Close    float64 `json:"close,omitempty"`
	Datetime vTime   `json:"datetime,omitempty"`
	High     float64 `json:"high,omitempty"`
	Low      float64 `json:"low,omitempty"`
	Open     float64 `json:"open,omitempty"`
	Volume   float64 `json:"volume,omitempty"`
}

type vTime time.Time

func (v *vTime) UnmarshalJSON(bs []byte) (err error) {
	millis, err := strconv.ParseInt(string(bs), 10, 64)
	if err != nil {
		return err
	}

	*v = vTime(time.Unix(0, millis*int64(time.Millisecond)))

	return nil
}

func (v *vTime) MarshalJSON() ([]byte, error) {
	t := time.Time(*v)
	millis := t.UnixNano() / int64(time.Millisecond)
	s := strconv.FormatInt(millis, 10)

	return []byte(s), nil
}

func (v *vTime) ToTime() time.Time {
	return time.Time(*v)
}

// PriceHistory https://developer.tdameritrade.com/price-history/apis
type PriceHistory struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Candles []Candle `json:"candles,omitempty"`
	Empty   bool     `json:"empty,omitempty"`
	Symbol  string   `json:"symbol,omitempty"` //"string"
}
