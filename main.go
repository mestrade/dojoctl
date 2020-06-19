package main

import (
	"dojoctl/pkg/dojo"
	"flag"
	"fmt"
	"os"
)

var (
	testURL   = "https://defectdojo.mydomain.com/api/v2"
	testToken = "yourtokenhere"
)

func main() {

	dojoSetupFile := flag.String("setup", "", "specify another defectdojo setup")

	productCommand := flag.NewFlagSet("product", flag.ExitOnError)
	productList := productCommand.Bool("list", false, "list all product")
	productRead := productCommand.String("read", "", "read a specific product product")
	productSet := productCommand.String("set", "", "read a specific product product")

	engagementCommand := flag.NewFlagSet("engagement", flag.ExitOnError)
	engagementList := engagementCommand.Bool("list", false, "List engagement in the current product")
	engagementProduct := engagementCommand.String("product", "", "where to search for engagement")
	engagementSet := engagementCommand.String("set", "", "set engagement")

	testCommand := flag.NewFlagSet("test", flag.ExitOnError)
	testList := testCommand.Bool("list", false, "list all tests in an engagement")
	testSet := testCommand.String("set", "", "set current test")

	uploadCommand := flag.NewFlagSet("upload", flag.ExitOnError)
	uploadFile := uploadCommand.String("file", "", "filename")
	uploadType := uploadCommand.String("type", "", "Report type")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "product":
		productCommand.Parse(os.Args[2:])
	case "engagement":
		engagementCommand.Parse(os.Args[2:])
	case "test":
		testCommand.Parse(os.Args[2:])
	case "upload":
		uploadCommand.Parse(os.Args[2:])
	}

	dj, err := dojo.NewDojoCtx(*dojoSetupFile)
	if err != nil {
		panic(err)
	}

	// List products
	if *productList {
		products, err := dj.ProductList()
		if err != nil {
			panic(err)
		}

		for _, product := range products {
			product.DisplayShort()
		}

		os.Exit(0)
	}

	// List products
	if len(*productRead) > 0 {
		product, err := dj.ProductByName(*productRead)
		if err != nil {
			panic(err)
		}

		product.DisplayShort()

		os.Exit(0)
	}

	if len(*productSet) > 0 {
		err := dj.SetProductByName(*productSet)
		if err != nil {
			fmt.Printf("Cannot set product: %v\n", err)
			os.Exit(-1)
		}
		err = dj.Save()
		if err != nil {
			fmt.Printf("Cannot save context: %v\n", err)
			os.Exit(-1)
		}
	}

	if *engagementList {

		if len(*engagementProduct) > 0 {
			dj.SetProductByName(*engagementProduct)
		}

		if len(dj.Context.CurrentProduct) == 0 {

			fmt.Printf("Need a product name")
			os.Exit(-1)
		}

		fmt.Printf("Look engagements for product %s\n", dj.Context.CurrentProduct)

		engagements, err := dj.EngagementList()
		if err != nil {
			panic(err)
		}

		for _, engagement := range engagements {
			engagement.DisplayShort()
		}

		os.Exit(0)
	}

	if len(*engagementSet) > 0 {
		err := dj.SetEngagementByName(*engagementSet)
		if err != nil {
			fmt.Printf("Cannot set engagement: %v\n", err)
			os.Exit(-1)
		}
		err = dj.Save()
		if err != nil {
			fmt.Printf("Cannot save context: %v\n", err)
			os.Exit(-1)
		}
	}

	// List tests
	if *testList {
		tests, err := dj.TestList()
		if err != nil {
			panic(err)
		}

		for _, test := range tests {
			test.DisplayShort()
		}

		os.Exit(0)
	}

	if len(*testSet) > 0 {
		err := dj.SetTestByName(*testSet)
		if err != nil {
			fmt.Printf("Cannot set test: %v\n", err)
			os.Exit(-1)
		}
		err = dj.Save()
		if err != nil {
			fmt.Printf("Cannot save test: %v\n", err)
			os.Exit(-1)
		}
	}

	if len(*uploadFile) > 0 && len(*uploadType) > 0 {

		err := dj.Upload(*uploadFile, *uploadType)
		if err != nil {
			fmt.Printf("Cannot upload report: %v\n", err)
			os.Exit(-1)
		}

	}

}
