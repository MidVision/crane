package subcommand

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"encoding/xml"
	"text/tabwriter"
)

func (cli *CraneSubcommand) Login(username, password, url *string) {

	cli.username = *username
	cli.password = *password
	cli.Url = *url

	fmt.Printf("\nTrying to log in as user '%s' to '%s'...\n", *username, *url)

	// CREATE BODY CONTENT FOR REQUEST
	type EnvelopeBody struct {
		Authentication string
	}
	
	// CREATE RESPONSE STRUCTURE
	type AuthenticationResponse struct {
		Result string
	}

	type RespBody struct {
		AuthenticationResponse AuthenticationResponse
		Fault Fault
	}
	
	type RespEnvelope struct {
		Body RespBody
	}
	
	respEnvelope := new(RespEnvelope)
	
	// CALL WEB SERVICE AND	DECODE RESPONSE
	resData, statusCode, err := cli.call("POST", EnvelopeBody{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if statusCode != 200 {
		fmt.Printf("Unable to connect to server '%s'.\n\n", cli.Url)
		fmt.Printf("Server returned response code %v: %v\n\n", statusCode, http.StatusText(statusCode))
		fmt.Println("Please check the credentials.\n")
		os.Exit(1)
	}

	err = xml.Unmarshal(resData, &respEnvelope)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// PRINT TABLE
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '*', 0)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "\t\t\n")
	if strings.Contains(respEnvelope.Body.AuthenticationResponse.Result, cli.username) {
		fmt.Fprintf(w, "\t Login successful as '%s' to '%s' \t\n", cli.username, cli.Url)
	} else {
		fmt.Fprintf(w, "\t %s \t\n", respEnvelope.Body.Fault.Faultstring)
	}
	fmt.Fprintf(w, "\t\t\n")
	fmt.Fprintln(w)
	w.Flush()
	
	// SAVE LOGIN CONFIGURATION
	if err := cli.saveLoginFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (cli *CraneSubcommand) Logout() {
	if err := cli.removeLoginFile(); err != nil {
		fmt.Println("\nWARNING: NO LOGIN SESSION FOUND!\n\nPlease, perform a new login before requesting any action.\n")
	} else {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 1, '*', 0)
		fmt.Fprintln(w)
		fmt.Fprintf(w, "\t\t\n")
		fmt.Fprintf(w, "\t Logged out successfully! \t\n")
		fmt.Fprintf(w, "\t\t\n")
		fmt.Fprintln(w)
		w.Flush()
	}
}