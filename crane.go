/* 
TODO:
 - Login credentials
     * Save in a hidden file (.crane in HOME directory)
     * Manage properties
	 * Manage files
 - Try to modularize the tool
 - Template for help
 - Template for usage
 - Try to output the progress and some output sentences
     * Implement the verbose/debug option
 - Controlar respuestas vacías y no mostrar nada
 - Mostrar un login interactivo?
*/

package main

import (
	//"os"
	"fmt"
	// More than logging it's output
	//"log"
	//"bytes"
	//"net/http"
	//"io/ioutil"
	//"text/tabwriter"
	"encoding/xml"
	"github.com/spf13/cobra"
	"github.com/MidVision/crane/subcommand"
)

type (
	Authentication struct {
		Username string `xml:"username"`
		Password string `xml:"password"`
	}	

	EnvelopeHeader struct {
		Credentials Authentication `xml:"get:authentication"`
	}		

	EnvelopeBody struct {
		Payload interface{}		`xml:"get:ListImages"`
	}

	Envelope struct {
		XMLName	xml.Name			`xml:"soapenv:Envelope"`
		SoapEnv		string			`xml:"xmlns:soapenv,attr"`
		CraneEnv	string			`xml:"xmlns:get,attr"`
		Header		EnvelopeHeader	`xml:"soapenv:Header"`
		Body		EnvelopeBody	`xml:"soapenv:Body"`
	}
)

var (
	//CLI
	cli *subcommand.CraneSubcommand

	//SOAP Envelope
	envelope *Envelope
	
	//Subcommand
	loginCmd  *cobra.Command
	listImagesCmd  *cobra.Command
	showImageCmd  *cobra.Command
	listContainersCmd  *cobra.Command
	showContainerCmd  *cobra.Command
	listHarborsCmd  *cobra.Command
	showHarborCmd  *cobra.Command
	listClustersCmd  *cobra.Command
	showClusterCmd  *cobra.Command
	listEnginesCmd  *cobra.Command
	showEngineCmd  *cobra.Command
	listShippingLanesCmd  *cobra.Command
	showShippingLaneCmd  *cobra.Command
	listVesselsCmd  *cobra.Command
	showVesselCmd  *cobra.Command
	deployCmd  *cobra.Command
	stopCmd  *cobra.Command
	startCmd  *cobra.Command
	restartCmd  *cobra.Command
	resizeCmd  *cobra.Command
	
	//Flags
	version bool
	debug bool = false
	
	username string
	password string
	url string	
)

func init() {
	envelope = new(Envelope)
	envelope.SoapEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	envelope.CraneEnv = "http://getcrane.com/"
	envelope.Header = EnvelopeHeader{}
	envelope.Header.Credentials = Authentication{}
	
	cli = new(subcommand.CraneSubcommand)
}

func main() {

	//login
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Performs a login to the Harbormaster server.",
		Long:  "This command performs a login to the Harbormaster server.\n\nOnce logged in you won't need to run this command again unless you want to connect to a different server, the credentials are kept during the session.\n\nRunning this command again with different parameters will result in a new login to a different server losing the previous session.",
		Run:	func(cmd *cobra.Command, args []string) {
			if username != "" && password != "" && url != "" {
				cli.Login(&username, &password, &url)
			} else {
				fmt.Println("WARNING: You are missing some parameters.\n")
				cmd.Help()
			}
		},
	}

	// TODO: remove the default values
	loginCmd.Flags().StringVarP(&username, "username", "", "", "User name to connect to Harbormaster.")
	loginCmd.Flags().StringVarP(&password, "password", "", "", "Password to connect to Harbormaster.")
	loginCmd.Flags().StringVarP(&url, "url", "", "", "URL to Harbormaster.")

	//listImages
	listImagesCmd = &cobra.Command{
		Use:   "listImages",
		Short: "Shows the Repository Overview of Harbormaster.",
		Long:  "This command shows all the Images contained in Harbormaster.",
		Run:	func(cmd *cobra.Command, args []string) {
			cli.ListImages()
		},
	}
	
	//showImage
	showImageCmd = &cobra.Command{
		Use:   "showImage",
		Short: "showImage",
		Long:  "showImage",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("showImage")
		},
	}
	
	//listContainers
	listContainersCmd = &cobra.Command{
		Use:   "listContainers",
		Short: "listContainers",
		Long:  "listContainers",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listContainers")
		},
	}
	
	//showContainer
	showContainerCmd = &cobra.Command{
		Use:   "showContainer",
		Short: "showContainer",
		Long:  "showContainer",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("showContainer")
		},
	}
	
	//listHarbors
	listHarborsCmd = &cobra.Command{
		Use:   "listHarbors",
		Short: "listHarbors",
		Long:  "listHarbors",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listHarbors")
		},
	}
	
	//showHarbor
	showHarborCmd = &cobra.Command{
		Use:   "listImages",
		Short: "listImages",
		Long:  "listImages",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listImages")
		},
	}
	
	//listClusters
	listClustersCmd = &cobra.Command{
		Use:   "listClusters",
		Short: "listClusters",
		Long:  "listClusters",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listClusters")
		},
	}
	
	//showCluster
	showClusterCmd = &cobra.Command{
		Use:   "showCluster",
		Short: "showCluster",
		Long:  "showCluster",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("showCluster")
		},
	}
	
	//listEngines
	listEnginesCmd = &cobra.Command{
		Use:   "listEngines",
		Short: "listEngines",
		Long:  "listEngines",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listEngines")
		},
	}
	
	//showEngine
	showEngineCmd = &cobra.Command{
		Use:   "showEngine",
		Short: "showEngine",
		Long:  "showEngine",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("showEngine")
		},
	}
	
	//listShippingLanes
	listShippingLanesCmd = &cobra.Command{
		Use:   "listShippingLanes",
		Short: "listShippingLanes",
		Long:  "listShippingLanes",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listShippingLanes")
		},
	}
	
	//showShippingLane
	showShippingLaneCmd = &cobra.Command{
		Use:   "showShippingLane",
		Short: "showShippingLane",
		Long:  "showShippingLane",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("showShippingLane")
		},
	}
	
	//listVessels
	listVesselsCmd = &cobra.Command{
		Use:   "listVessels",
		Short: "listVessels",
		Long:  "listVessels",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("listVessels")
		},
	}
	
	//showVessel
	showVesselCmd = &cobra.Command{
		Use:   "showVessel",
		Short: "showVessel",
		Long:  "showVessel",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("showVessel")
		},
	}
	
	//deploy
	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "deploy",
		Long:  "deploy",
		Run:	func(cmd *cobra.Command, args []string) {
			/*
			fmt.Println(resp)
			
			if err != nil {
				fmt.Println(err)
			}
			*/
			fmt.Println("deploy")
		},
	}
	
	//stop
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop",
		Long:  "stop",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("stop")
		},
	}
	
	//start
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start",
		Long:  "start",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("start")
		},
	}
	
	//restart
	restartCmd = &cobra.Command{
		Use:   "restart",
		Short: "restart",
		Long:  "restart",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("restart")
		},
	}
	
	//resize
	resizeCmd = &cobra.Command{
		Use:   "resize",
		Short: "resize",
		Long:  "resize",
		Run:	func(cmd *cobra.Command, args []string) {
			fmt.Println("resize")
		},
	}
	
	// MAIN
	var craneCmd = &cobra.Command {
		Use:	"crane",
		Short:	"Crane is a command line tool for Harbormaster.",
		Long:	"Crane is a command line tool for Harbormaster.\n\nUse this command to interact with the features provided by Harbormaster.",
		Run:	func(cmd *cobra.Command, args []string) {
			if (version) {
				fmt.Println("Crane - The command line interface tool for Harbormaster - v0.1")
			} else {
				cmd.Help()
			}
		},
	}
	
	craneCmd.Flags().BoolVarP(&version, "version", "v", false, "Shows the version of 'crane'.")
	
	craneCmd.AddCommand(loginCmd)
	craneCmd.AddCommand(listImagesCmd)
	craneCmd.AddCommand(showImageCmd)
	craneCmd.AddCommand(listContainersCmd)
	craneCmd.AddCommand(showContainerCmd)
	craneCmd.AddCommand(listHarborsCmd)
	craneCmd.AddCommand(showHarborCmd)
	craneCmd.AddCommand(listClustersCmd)
	craneCmd.AddCommand(showClusterCmd)
	craneCmd.AddCommand(listEnginesCmd)
	craneCmd.AddCommand(showEngineCmd)
	craneCmd.AddCommand(listShippingLanesCmd)
	craneCmd.AddCommand(showShippingLaneCmd)
	craneCmd.AddCommand(listVesselsCmd)
	craneCmd.AddCommand(showVesselCmd)
	craneCmd.AddCommand(deployCmd)
	craneCmd.AddCommand(stopCmd)
	craneCmd.AddCommand(startCmd)
	craneCmd.AddCommand(restartCmd)
	craneCmd.AddCommand(resizeCmd)
	
	craneCmd.Execute()
}

