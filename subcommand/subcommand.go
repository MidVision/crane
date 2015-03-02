package subcommand

import (
	"os"
	"fmt"
	"path"
	"bytes"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
)

const (
	LOGINFILE = ".crane"
	WSRELPATH = "/ws/HarborMaster"
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
		Url 		string	`json:"url"`
		Token		string	`json:"token"`
		username	string	`json:",omitempty"`
		password	string	`json:",omitempty"`
	}
)

var (
	envelope *Envelope

	// For debugging purposes only
	debug bool = false
)

func init() {
	envelope = new(Envelope)
	envelope.SoapEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	envelope.CraneEnv = "http://getcrane.com/"
	envelope.Header = EnvelopeHeader{}
	envelope.Header.Credentials = Authentication{}
}

func (cli *CraneSubcommand) loadLoginFile() error {
	loginFilePath := path.Join(getHome(), LOGINFILE)
	
	if _, err := os.Stat(loginFilePath); err != nil {
		return fmt.Errorf("\nERROR: NO LOGIN FOUND!\n\nPlease, perform a login before requesting any action.\n")
	}

	content, err := ioutil.ReadFile(loginFilePath)
	if err != nil {
		return fmt.Errorf("\nERROR: INVALID LOGIN FOUND!\n\nPlease, perform a new login before requesting any action.\n")
	}

	if err := json.Unmarshal(content, cli); err != nil {
		return fmt.Errorf("\nERROR: INVALID LOGIN FOUND!\n\nPlease, perform a new login before requesting any action.\n")
	} else {
		cli.username, cli.password, err = decodeToken(cli.Token)
		cli.Token = ""
		if err != nil {
			return fmt.Errorf("\nERROR: INVALID LOGIN FOUND!\n\nPlease, perform a new login before requesting any action.\n")
		}
	}
	return nil
}

func (cli *CraneSubcommand) saveLoginFile() error {
	loginFilePath := path.Join(getHome(), LOGINFILE)
	
	cli.Token = encodeToken(cli.username, cli.password)
	cli.username = ""
	cli.password = ""

	content, err := json.MarshalIndent(cli, "", "\t")
	if err != nil {
		return err
	} else {
		err = ioutil.WriteFile(loginFilePath, content, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cli *CraneSubcommand) removeLoginFile() error {
	loginFilePath := path.Join(getHome(), LOGINFILE)
	return os.Remove(loginFilePath)
}

func (cli *CraneSubcommand) call(method string, bodyContent interface{}) ([]byte, int, error) {

	// SET CREDENTIALS
	envelope.Header.Credentials.Username = cli.username
	envelope.Header.Credentials.Password = cli.password

	// SET BODY CONTENT
	envelope.Body = bodyContent

	// ENCODE REQUEST
	reqData, err := xml.MarshalIndent(envelope, "  ", "    ")
	if err != nil {
		return nil, -1, err
	}

	if debug {
		fmt.Printf("[DEBUG] Request:\n\n%v\n\n", string(reqData))
	}

	// SETUP CALL
	client := http.Client{}
	req, err := http.NewRequest(method, WSRELPATH, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, -1, fmt.Errorf("Error creating request: \n\n%v\n", err)
	}
	u, err := url.Parse(cli.Url)
	if err != nil {
		return nil, -1, fmt.Errorf("Error parsing URL provided '%s': \n\n%v\n", cli.Url, err)
	}
	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme

	req.Header.Add("Content-Type", "text/xml")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, fmt.Errorf("Error trying to connect to '%s': \n\n%v\n", cli.Url, err)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	if debug {
		fmt.Printf("[DEBUG] Response:\n\n%v\n\n", string(resData))
	}

	resp.Body.Close()

	return resData, resp.StatusCode, nil
}