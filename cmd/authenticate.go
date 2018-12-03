package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/provideservices/provide-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Authenticate using your developer credentials and receive a valid API token",
	Long: `Authenticate using user credentials retrieved from provide.services and receive a
valid API token which can be used to access the networks and application APIs.`,
	Run: authenticate,
}

func init() {
	rootCmd.AddCommand(authenticateCmd)
}

func authenticate(cmd *cobra.Command, args []string) {
	email := doEmailPrompt()
	passwd := doPasswordPrompt()

	status, resp, err := provide.Authenticate(email, passwd)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if status != 201 {
		log.Println("Authentication failed")
		os.Exit(1)
	}

	if token, tokenOk := resp.(map[string]interface{})["token"].(map[string]interface{}); tokenOk {
		if apiToken, apiTokenOk := token["token"].(string); apiTokenOk {
			cacheAPIToken(apiToken)
			log.Printf("Authentication successful")
		}
	}
}

func doEmailPrompt() string {
	fmt.Print("Email: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	email = strings.TrimSpace(email)
	if email == "" {
		log.Println("Failed to read email from stdin")
		os.Exit(1)
	}
	// Remove both \r and \n as in windows uses LF and CF chars
	return strings.Replace(email, "\r\n", "", -1)
}

func doPasswordPrompt() string {
	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	password := string(bytePassword)
	passwd := strings.TrimSpace(password)
	if passwd == "" {
		log.Println("Failed to read password from stdin")
		os.Exit(1)
	}
	// Remove both \r and \n as in windows uses LF and CF chars
	return strings.Replace(passwd, "\r\n", "", -1)
}

func cacheAPIToken(token string) {
	viper.Set(authTokenConfigKey, token)
	viper.WriteConfig()
}
