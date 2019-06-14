package torn

import (
	"io/ioutil"
	"net/http"
)

// Type definition for parameter usage
type endpoint string

// API endpoint constants
const (
	base          endpoint = "https://api.torn.com/"
	apiTorn       endpoint = base + "torn/"
	apiMarket     endpoint = base + "market/"
	apiCompany    endpoint = base + "company/"
	apiFaction    endpoint = base + "faction/"
	apiProperties endpoint = base + "properties/"
	apiUser       endpoint = base + "user/"
)

// Session represents
// a state for given API operations.
// This includes a Torn users API key
// to access the Torn API, as well as
// stateful data, such as item prices.
type Session struct {
	apiKey string
}

// NewSession returns a *Session
// initialized with the give apiKey.
func NewSession(apiKey string) *Session {
	return &Session{apiKey: apiKey}
}

func (s *Session) buildEndpoint(api endpoint, args map[string]string) string {
	ep := string(api) + "?"
	for k, v := range args {
		ep += k + "=" + v + "&"
	}
	ep += "key=" + s.apiKey
	return ep
}

func (s *Session) callAPI(api endpoint, args map[string]string) (data []byte, err error) {
	url := s.buildEndpoint(api, args)
	resp, err := http.Get(url)
	if err == nil {
		data, err = ioutil.ReadAll(resp.Body)
	}
	return
}
