package dojo

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

var (
	defaultFilename = "config"
)

type Setup struct {
	ApiBaseUrl string `toml:"api_base_url"`
	Token      string `toml:"token"`
}

type Context struct {
	CurrentProduct    string `toml:"product"`
	CurrentEngagement string `toml:"engagement"`
	CurrentTest       string `toml:"test"`

	// Local info retrieved from dojo
	currentProductID    int
	currentEngagementID int
	currentTestID       int
}

type Ctx struct {
	Filename string `toml:"-"`
	Setup    Setup
	Context  Context
	Debug    bool
}

func NewDojoCtx(filename string) (*Ctx, error) {

	var setupFile string

	if len(filename) > 0 {
		setupFile = filename
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		setupFile = fmt.Sprintf("%s/.dojoctl/%s", home, defaultFilename)
	}

	data, err := ioutil.ReadFile(setupFile)
	if err != nil {
		return nil, err
	}

	var setup Ctx
	err = toml.Unmarshal(data, &setup)
	if err != nil {
		return nil, err
	}

	// save setup file to save context later
	setup.Filename = setupFile

	// if the context contain a product name - look for product id
	if len(setup.Context.CurrentProduct) > 0 {
		product, err := setup.ProductByName(setup.Context.CurrentProduct)
		if err != nil {
			return nil, err
		}

		setup.Context.currentProductID = product.Id
	}

	// if the context contain an engagement name - look for engagement id
	if len(setup.Context.CurrentEngagement) > 0 {
		eng, err := setup.EngagementByName(setup.Context.CurrentEngagement)
		if err != nil {
			return nil, err
		}

		setup.Context.currentEngagementID = eng.Id
	}

	return &setup, nil
}

func (ctx *Ctx) Ping() error {
	return nil
}

func (ctx *Ctx) SetProductByName(name string) error {
	product, err := ctx.ProductByName(name)
	if err != nil {
		return err
	}

	ctx.Context.CurrentProduct = product.Name
	ctx.Context.currentProductID = product.Id

	// reset engagement and tests
	ctx.Context.CurrentEngagement = ""
	ctx.Context.currentEngagementID = 0
	ctx.Context.CurrentTest = ""
	ctx.Context.currentTestID = 0
	return nil
}

func (ctx *Ctx) SetEngagementByName(name string) error {
	eng, err := ctx.EngagementByName(name)
	if err != nil {
		return err
	}
	ctx.Context.CurrentEngagement = eng.Name
	ctx.Context.currentEngagementID = eng.Id

	// reset test
	ctx.Context.CurrentTest = ""
	ctx.Context.currentTestID = 0

	return nil
}

func (ctx *Ctx) SetTestByName(name string) error {
	test, err := ctx.TestByName(name)
	if err != nil {
		return err
	}
	ctx.Context.CurrentTest = test.Title
	ctx.Context.currentTestID = test.Id
	return nil
}

func (ctx *Ctx) Save() error {

	data, err := toml.Marshal(ctx)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ctx.Filename, data, 0600)
	if err != nil {
		return err
	}

	fmt.Printf("Context saved in file: %s\n", ctx.Filename)

	return nil
}
