package subcommand

import (
	"os"
	"fmt"
	"bytes"
	"runtime"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type (
	// SOAP Request
	Authentication struct {
		Username	string	`xml:"username"`
		Password	string	`xml:"password"`
	}	

	EnvelopeHeader struct {
		Credentials Authentication 	`xml:"get:authentication"`
	}		

	Envelope struct {
		XMLName	xml.Name			`xml:"soapenv:Envelope"`
		SoapEnv		string			`xml:"xmlns:soapenv,attr"`
		CraneEnv	string			`xml:"xmlns:get,attr"`
		Header		EnvelopeHeader	`xml:"soapenv:Header"`
		Body		interface{}		`xml:"soapenv:Body"`
	}

	// SOAP Response
	Fault struct {
		Faultcode 	string `xml:"faultcode"`
		Faultstring string `xml:"faultstring"`
	}
	
	RespEnvelope interface{}

	// CLI
	CraneSubcommand struct {
		url 	string
		auth 	string
		// TODO: remove
		username string
		password string
	}
)

var (
	envelope *Envelope
	
	debug bool = true
)

func init() {
	envelope = new(Envelope)
	envelope.SoapEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	envelope.CraneEnv = "http://getcrane.com/"
	envelope.Header = EnvelopeHeader{}
	envelope.Header.Credentials = Authentication{}
}

func (cli *CraneSubcommand) LoadLoginFile() error {
	return nil
}

func (cli *CraneSubcommand) SaveLoginFile() error {
	return nil
}

func (cli *CraneSubcommand) call() (string, error) {

	envelope.Header.Credentials.Username = cli.username
	envelope.Header.Credentials.Password = cli.password
	
	type EnvelopeBody struct {
		Authentication string
	}
	
	envelope.Body = EnvelopeBody{}
	
	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	encoder.Indent("  ", "    ")
	err := encoder.Encode(envelope)
	if err != nil {
		fmt.Println("Could not encode request")
	}
	
	if debug {
		fmt.Printf("[DEBUG] Request:\n\n%v\n\n", buffer)
	}
	
	client := http.Client{}
	req, err := http.NewRequest("POST", cli.url, buffer)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-Type", "text/xml")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(resp.Status)
	}
	in := string(b)

	if debug {
		fmt.Printf("[DEBUG] Response:\n\n%v\n\n", in)
	}
	return in, nil
}

func getHome() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}