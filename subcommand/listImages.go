package subcommand

import (
	"os"
	"fmt"
	"net/http"
	"text/tabwriter"
	"encoding/xml"
)

func (cli *CraneSubcommand) ListImages() {

	// LOAD LOGIN CONFIGURATION
	if err := cli.loadLoginFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Printing the list of available container images...\n")

	// CREATE BODY CONTENT FOR REQUEST
	type EnvelopeBody struct {
		ListImages string
	}

	// CREATE RESPONSE STRUCTURE
	type ImageTag struct {
		TagName string
		Approved bool
	}

	type Image struct {
		Username string
		Password string
		ImageSource string
		RepositoryUrl string
		EmailAddress string
		ApprovalAllImages bool
		ImageName string
		Tags []ImageTag `xml:"ImageVersion_Image>ImageTag"`
	}

	type ListImagesResponse struct {
		Images []Image `xml:"Image"`
	}

	type RespBody struct {
		ListImagesResponse ListImagesResponse
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
		os.Exit(1)
	}

	err = xml.Unmarshal(resData, &respEnvelope)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// PRINT TABLE
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 12, 8, 1, ' ', 0)
	// Table header
	fmt.Fprintln(w, "NAME\tSOURCE\tEMAIL ADDRESS\tTAG")
	fmt.Fprintln(w, "\t\t\t\t")
	for _, image := range respEnvelope.Body.ListImagesResponse.Images {
		for _, tag := range image.Tags {
			if tag.Approved {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\tApproved\n", image.ImageName, image.ImageSource, image.EmailAddress, tag.TagName)
			} else {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\tNot approved\n", image.ImageName, image.ImageSource, image.EmailAddress, tag.TagName)
			}
		}
	}
	fmt.Fprintln(w)
	w.Flush()
}