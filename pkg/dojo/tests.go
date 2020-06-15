package dojo

import "fmt"

var (
	testListCall = "/tests/"
)

type Test struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

type testResponse struct {
	Count int    `json:"count"`
	List  []Test `json:"results"`
}

func (ctx *Ctx) TestList() ([]Test, error) {

	url := fmt.Sprintf("%s%s?product=%d&engagement=%d", ctx.Setup.ApiBaseUrl, testListCall, ctx.Context.currentProductID, ctx.Context.currentEngagementID)

	var tests testResponse
	err := ctx.req("GET", url, &tests)
	if err != nil {
		return nil, err
	}

	return tests.List, nil
}

func (ctx *Ctx) TestByName(name string) (*Test, error) {

	url := fmt.Sprintf("%s%s?name=%s&product=%d&engagement=%d", ctx.Setup.ApiBaseUrl, testListCall, name, ctx.Context.currentProductID, ctx.Context.currentEngagementID)

	var tests testResponse
	err := ctx.req("GET", url, &tests)
	if err != nil {
		return nil, err
	}

	if len(tests.List) == 0 {
		return nil, fmt.Errorf("No such test")
	}

	return &tests.List[0], nil
}

func (t *Test) DisplayShort() {
	fmt.Printf("Test: %s (id: %d)\n", t.Title, t.Id)
}
