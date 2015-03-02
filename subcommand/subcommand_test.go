package subcommand

import (
	"testing"
)

func TestLoginFile(t *testing.T) {
	url := "http://localhost:8080"
	username := "harborUser"
	password := "harborPass"
	
	cli := &CraneSubcommand{Url: url, username: username, password: password}
	t.Logf("CLI set to: %+v", cli)
	if err := cli.saveLoginFile(); err != nil {
		t.Error(err)
	}
	t.Logf("File saved with no errors: %+v", cli)
	cli.Url = ""
	cli.Token = ""
	cli.username = ""
	cli.password = ""
	t.Logf("CLI re-set to: %+v", cli)

	if err := cli.loadLoginFile(); err != nil && cli.Url == url && cli.username == username && cli.password == password {
			t.Error(err)
	}
	t.Logf("File loaded with no errors: %+v", cli)
	
	if err := cli.removeLoginFile(); err != nil {
		t.Error(err)
	}
	t.Log("File removed with no errors.")
	
	if err := cli.loadLoginFile(); err != nil {
		t.Logf("Not able to load: %v", err)
	}
}

func TestLogin(t *testing.T) {
	
}