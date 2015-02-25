package subcommand

import (
	"os"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"text/tabwriter"
	"encoding/xml"
)

func (subcommand *CraneSubcommand) ListImages() {
	fmt.Println("Printing the list of available container images...\n")
	
	// TODO: remove
	url := "http://localhost:8080/ws/HarborMaster"
	
	envelope.Header.Credentials.Username = "harborWS"
	envelope.Header.Credentials.Password = "123456Ab"
	
	type EnvelopeBody struct {
		ListImages string
	}
	
	body := EnvelopeBody{}

	envelope.Body = body
	
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
	req, err := http.NewRequest("POST", url, buffer)
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

	parser := xml.NewDecoder(bytes.NewBufferString(in))
	
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
	
	parser.Decode(respEnvelope)
	
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
	
	resp.Body.Close()
}