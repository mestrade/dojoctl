package dojo

import "fmt"

var (
	productListCall = "/products/"
)

type Product struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type productResponse struct {
	Count int       `json:"count"`
	List  []Product `json:"results"`
}

func (ctx *Ctx) ProductList() ([]Product, error) {

	url := fmt.Sprintf("%s%s", ctx.Setup.ApiBaseUrl, productListCall)

	var products productResponse
	err := ctx.req("GET", url, &products)
	if err != nil {
		return nil, err
	}

	return products.List, nil
}

func (ctx *Ctx) ProductByName(name string) (*Product, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("Need a valid project name")
	}

	url := fmt.Sprintf("%s%s?name=%s", ctx.Setup.ApiBaseUrl, productListCall, name)

	var products productResponse
	err := ctx.req("GET", url, &products)
	if err != nil {
		return nil, err
	}

	if len(products.List) == 0 {
		return nil, fmt.Errorf("cannot find such product")
	}

	if len(products.List) > 1 {
		return nil, fmt.Errorf("Found multiple product with this name")
	}

	return &products.List[0], nil
}

func (p *Product) DisplayShort() {
	fmt.Printf("Product: %s (id: %d)\n", p.Name, p.Id)
}
