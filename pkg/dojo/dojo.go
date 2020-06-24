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

func (ctx *Ctx) RetrieveCurrentProductID() error {
	// if the context contain a product name - look for product id
	if len(ctx.Context.CurrentProduct) > 0 {
		product, err := ctx.ProductByName(ctx.Context.CurrentProduct)
		if err != nil {
			return fmt.Errorf("unable to find product: %v", err)
		}

		ctx.Context.currentProductID = product.Id
		return nil
	}

	return fmt.Errorf("Need a valid product name")
}

func (ctx *Ctx) RetrieveCurrentEngagementID() error {
	// if the context contain a product name - look for product id
	if len(ctx.Context.CurrentEngagement) > 0 {
		eng, err := ctx.EngagementByName(ctx.Context.CurrentEngagement)
		if err != nil {
			return fmt.Errorf("unable to find engagement: %v", err)
		}

		ctx.Context.currentEngagementID = eng.Id
		return nil
	}

	return fmt.Errorf("Need a valid engagement name")
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
		fmt.Printf("No config file found - setup context to %s\n", setupFile)
		new := &Ctx{Filename: setupFile}
		return new, nil
	}

	var setup Ctx
	err = toml.Unmarshal(data, &setup)
	if err != nil {
		return nil, err
	}

	// save setup file to save context later
	setup.Filename = setupFile

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
		fmt.Printf("Unable to save setup: %v\n", err)
		return err
	}

	fmt.Printf("Context saved in file: %s\n", ctx.Filename)

	return nil
}
