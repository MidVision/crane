package subcommand

import (
	"os"
	"fmt"
	"bytes"
	"strings"
	"text/tabwriter"
	"encoding/xml"
)

func (cli *CraneSubcommand) Login(username, password, url *string) {
	fmt.Println("Logging in...\n")
	
	*url = *url + "/ws/HarborMaster"
	
	cli.username = *username
	cli.password = *password
	cli.url = *url
	
	in, _ := cli.call()

	parser := xml.NewDecoder(bytes.NewBufferString(in))
	
	type AuthenticationResponse struct {
		Result string
	}

	type RespBody struct {
		AuthenticationResponse AuthenticationResponse
		Fault Fault
	}
	
	type RespEnvelope struct {
		Body 		RespBody
	}
	
	respEnvelope := new(RespEnvelope)
	
	parser.Decode(respEnvelope)
	
	//resp.Body.Close()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '*', 0)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "\t\t\n")
	if strings.Contains(respEnvelope.Body.AuthenticationResponse.Result, *username) {
		fmt.Fprintf(w, "\t Login successful as '%s' to '%s' \t\n", *username, *url + "/ws/HarborMaster")
	} else {
		fmt.Fprintf(w, "\t %s \t\n", respEnvelope.Body.Fault.Faultstring)
	}
	fmt.Fprintf(w, "\t\t\n")
	fmt.Fprintln(w)
	w.Flush()
}

/* 
{
	"https://index.docker.io/v1/":
		{
			"auth":"bWlkdmlzaW9uOm0xZHYxczFvbg==",
			"email":"support@midvision.com"
		}
} 
*/