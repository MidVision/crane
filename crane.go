// [_Command-line flags_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line
// programs. For example, in `wc -l` the `-l` is a
// command-line flag.

package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import "flag"
import "fmt"
//import "bufio"
//import "os"
//import "gopass"
import "net/http"
import "bytes"

func main() {

    //reader := bufio.NewReader(os.Stdin)
	
    // Basic flag declarations are available for string,
    // integer, and boolean options. Here we declare a
    // string flag `word` with a default value `"foo"`
    // and a short description. This `flag.String` function
    // returns a string pointer (not a string value);
    // we'll see how to use this pointer below.
    username := flag.String("username", "harbor", "The username used to log into Crane.")
    password := flag.String("password", "mypass", "The password used to log into Crane.")
    url := flag.String("url", "http://app.getcrane.com", "The URL used to log into Crane.")
    //auth := flag.String("auth", "unsetDefault", "Set the authentication credentials to log in to Crane")
    // This declares `numb` and `fork` flags, using a
    // similar approach to the `word` flag.
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // It's also possible to declare an option that uses an
    // existing var declared elsewhere in the program.
    // Note that we need to pass in a pointer to the flag
    // declaration function.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Once all flags are declared, call `flag.Parse()`
    // to execute the command-line parsing.
    flag.Parse()

    // Here we'll just dump out the parsed options and
    // any trailing positional arguments. Note that we
    // need to dereference the points with e.g. `*wordPtr`
    // to get the actual option values.
    fmt.Println("Username:", *username)
    fmt.Println("Password:", *password)
    fmt.Println("URL:", *url)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
	
	reader := bytes.NewBufferString("<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:exam=\"http://www.example.com/\"><soapenv:Header><exam:authentication><username>" + *username + "</username><password>" + *password + "</password></exam:authentication></soapenv:Header><soapenv:Body><exam:WS_StartVoyage><ShippingLaneName>aws-shipping-lane-rd-dev</ShippingLaneName></exam:WS_StartVoyage></soapenv:Body></soapenv:Envelope>")
	
	resp, err := http.Post(*url, "encoding/xml", reader)
	
	if err != nil {
		fmt.Println(resp)
		fmt.Println(err)
	}

	/*
    if *auth == "unsetDefault" {
        fmt.Print("Username:", )
        username, _ := reader.ReadString('\n')
        password, _ := gopass.GetPass("Password:")
         fmt.Print("URL:", )
        url, _ := reader.ReadString('\n')

        fmt.Println("Username:", username)
        fmt.Println("Password:", password)
        fmt.Println("Url:", url)
    }
	*/
}

