package dojo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	defaultTimeout = "30s"
)

func (ctx *Ctx) req(method, path string, object interface{}) error {

	timeout, err := time.ParseDuration(defaultTimeout)
	if err != nil {
		return fmt.Errorf("Wrong timeout value")
	}

	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: timeout,
			}).Dial,
			TLSHandshakeTimeout:   timeout,
			ResponseHeaderTimeout: timeout,
			ExpectContinueTimeout: timeout,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", ctx.Setup.Token))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if ctx.Debug {
		fmt.Printf("Req body: %s\n", string(body))
	}

	err = json.Unmarshal(body, &object)
	if err != nil {
		return err
	}

	return nil
}

func (ctx *Ctx) post(method, path string, filename string, reportType string) error {

	timeout, err := time.ParseDuration(defaultTimeout)
	if err != nil {
		return fmt.Errorf("Wrong timeout value")
	}

	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: timeout,
			}).Dial,
			TLSHandshakeTimeout:   timeout,
			ResponseHeaderTimeout: timeout,
			ExpectContinueTimeout: timeout,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// add engagement ID
	engagementIDStr := fmt.Sprintf("%d", ctx.Context.currentEngagementID)
	eng, err := writer.CreateFormField("engagement")
	eng.Write([]byte(engagementIDStr))

	// add report type
	rt, err := writer.CreateFormField("scan_type")
	rt.Write([]byte(reportType))

	// add report type
	now := time.Now().Format("2006-01-02")
	date, err := writer.CreateFormField("scan_date")
	date.Write([]byte(now))

	fmt.Printf("Open report: %s\n", filename)
	// Process report
	data, err := os.Open(filename)
	if err != nil {
		return err
	}

	part, err := writer.CreateFormFile("file", "file")
	if _, err := io.Copy(part, data); err != nil {
		return err
	}

	// end multipart
	err = writer.Close()
	if err != nil {
		return err
	}

	fmt.Printf("Will post %d bytes\n", body.Len())

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", ctx.Setup.Token))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {

		errBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("err body: %s\n", string(errBody))
		return fmt.Errorf("something wrong happened: http err code: %s", resp.Status)
	}

	return nil
}
