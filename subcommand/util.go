package subcommand

import (
	"os"
	"fmt"
	"strings"
	"runtime"
	"encoding/base64"
)

func encodeToken(username, password string) string {
	tokenStr := username + ":" + password
	msg := []byte(tokenStr)
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(msg)))
	base64.StdEncoding.Encode(encoded, msg)
	return string(encoded)
}

func decodeToken(tokenStr string) (string, string, error) {
	decLen := base64.StdEncoding.DecodedLen(len(tokenStr))
	decoded := make([]byte, decLen)
	tokenByte := []byte(tokenStr)
	n, err := base64.StdEncoding.Decode(decoded, tokenByte)
	if err != nil {
		return "", "", err
	}
	if n > decLen {
		return "", "", fmt.Errorf("Something went wrong decoding the authentication token.")
	}
	arr := strings.SplitN(string(decoded), ":", 2)
	if len(arr) != 2 {
		return "", "", fmt.Errorf("Invalid authentication token.")
	}
	username := arr[0]
	password := strings.Trim(arr[1], "\x00")
	return username, password, nil
}

func getHome() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}