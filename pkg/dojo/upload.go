package dojo

import "fmt"

var (
	uploadCall = "/import-scan/"
)

func (ctx *Ctx) Upload(filename, reportType string) error {

	url := fmt.Sprintf("%s%s", ctx.Setup.ApiBaseUrl, uploadCall)

	err := ctx.post("POST", url, filename, reportType)
	if err != nil {
		return err
	}

	return nil
}
