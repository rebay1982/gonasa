package gonasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rebay1982/gonasa/internal/app/api"
)

func (c *MarsAPIClient) MarsManifestAPI() (*api.Rovers, error) {

	res, err := http.Get(fmt.Sprintf("%s?api_key=%s", marsAPIBaseURL, c.apiKey))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ret api.Rovers
	if err := json.Unmarshal(bodyBytes, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
