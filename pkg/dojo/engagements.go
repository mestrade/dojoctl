package dojo

import "fmt"

var (
	engagementListCall = "/engagements/"
)

type Engagement struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type engagementResponse struct {
	Count int          `json:"count"`
	Next  string       `json:"next"`
	List  []Engagement `json:"results"`
}

func (ctx *Ctx) EngagementList() ([]Engagement, error) {

	var next string
	var out []Engagement

	for {

		var url string

		if len(next) > 0 {
			url = next
		} else {
			url = fmt.Sprintf("%s%s", ctx.Setup.ApiBaseUrl, engagementListCall)
		}

		if len(ctx.Context.CurrentProduct) > 0 {
			url = fmt.Sprintf("%s?product=%d", url, ctx.Context.currentProductID)
		}

		var engagements engagementResponse
		err := ctx.req("GET", url, &engagements)
		if err != nil {
			return nil, err
		}

		out = append(out, engagements.List...)

		if len(engagements.Next) > 0 {
			fmt.Printf("Got more engagements here: %s\n", engagements.Next)
			next = engagements.Next
		} else {
			break
		}
	}

	return out, nil
}

func (ctx *Ctx) EngagementByName(name string) (*Engagement, error) {

	url := fmt.Sprintf("%s%s?name=%s&product=%d", ctx.Setup.ApiBaseUrl, engagementListCall, name, ctx.Context.currentProductID)

	var engagements engagementResponse
	err := ctx.req("GET", url, &engagements)
	if err != nil {
		return nil, err
	}

	if len(engagements.List) == 0 {
		return nil, fmt.Errorf("No such engagement")
	}

	if len(engagements.List) > 1 {
		return nil, fmt.Errorf("Found multiple engagement with this name")
	}

	return &engagements.List[0], nil
}

func (p *Engagement) DisplayShort() {
	fmt.Printf("Engagement: %s\n", p.Name)
}
