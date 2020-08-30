package congress

import (
	"fmt"
	"net/http"
)

type committeeSuccessResponse struct {
	Status    string            `json:"status"`
	Copyright string            `json:"copyright"`
	Results   []committeeResult `json:"results"`
}

type committeeResult struct {
	Congress   string      `json:"congress"`
	Chamber    string      `json:"chamber"`
	NumResults int         `json:"num_results"`
	Committees []Committee `json:"committees"`
}

type Committee struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Chamber       string         `json:"chamber"`
	Url           string         `json:"url"`
	ApiUrl        string         `json:"api_url"`
	Chair         string         `json:"chair"`
	ChairId       string         `json:"chair_id"`
	ChairParty    string         `json:"chair_party"`
	ChairState    string         `json:"chair_state"`
	ChairUri      string         `json:"chair_uri"`
	RankingMember string         `json:"ranking_member"`
	Subcommittees []Subcommittee `json:"subcommittees"`
}

type Subcommittee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	ApiUri string `json:"api_uri"`
}

type CommitteeQueryParameters struct {
	Congress string
	Chamber  string
}

func (c *Client) GetCommittees(options *CommitteeQueryParameters) (*committeeSuccessResponse, error) {
	congress := "116"
	chamber := "senate"

	if options != nil {
		congress = options.Congress
		chamber = options.Chamber
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/committees.json", c.BaseURL, congress, chamber), nil)
	if err != nil {
		return nil, err
	}

	res := committeeSuccessResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
