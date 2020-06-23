package main

import (
	"dojoctl/cmd"
	"dojoctl/pkg/dojo"
	"fmt"
	"os"
)

func main() {

	var err error

	cmd.DojoCtx, err = dojo.NewDojoCtx("")
	if err != nil {
		fmt.Printf("Unable to connect to defectdojo: %v\n", err)
		os.Exit(-1)
	}

	cmd.Execute()

}
