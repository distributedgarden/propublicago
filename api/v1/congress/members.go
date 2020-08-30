package congress

import (
	"fmt"
	"net/http"
)

type memberSuccessResponse struct {
	Status    string         `json:"status"`
	Copyright string         `json:"copyright"`
	Results   []MemberResult `json:"results"`
}

type MemberResult struct {
	Congress   string   `json:"congress"`
	Chamber    string   `json:"chamber"`
	NumResults int      `json:"num_results"`
	Offset     int      `json:"offset"`
	Members    []Member `json:"members"`
}

type Member struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	ShortTitle      string `json:short_title"`
	ApiUri          string `json:"api_uri"`
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Suffix          string `json:"suffix"`
	DateOfBirth     string `json:"date_of_birth"`
	Gender          string `json:"gender"`
	Party           string `json:"party"`
	LeadershipRole  string `json:"leadership_role"`
	TwitterAccount  string `json:"twitter_account"`
	FacebookAccount string `json:"facebook_account"`
	YoutubeAccount  string `json:"youtube_account"`
	GovtrackId      string `json:"govtrack_id"`
	CspanId         string `json:"cspan_id"`
	State           string `json:"state"`
	NextElection    string `json:"next_election"`
}

type MembersQueryParameters struct {
	Congress string
	Chamber  string
}

func (c *Client) GetMembers(options *MembersQueryParameters) (*memberSuccessResponse, error) {
	congress := "116"
	chamber := "senate"

	if options != nil {
		congress = options.Congress
		chamber = options.Chamber
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/members.json", c.BaseURL, congress, chamber), nil)
	if err != nil {
		return nil, err
	}

	res := memberSuccessResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
