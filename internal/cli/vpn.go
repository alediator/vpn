package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const (
	hostFlag                  = "host"
	hostEnv                   = "VPN_HOST"
	userFlag                  = "user"
	userEnv                   = "VPN_USER"
	staticPasswordFlag        = "password"
	staticPasswordEnv         = "VPN_STATIC_PASSWORD"
	tokenKeyFlag              = "token"
	scriptFlag                = "script"
	scriptEnv                 = "VPN_POST_SCRIPT"
	openconnectAdditionalFlag = "flags"
	openconnectAdditionalEnv  = "VPN_ADDITIONAL_FLAGS"
)

// InitVpnCmd initialize beers command
func InitVpnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Helper to connect to a VPN through openconnect. \nNote that this is able to read `.env` files (by default in `/home/$USER/.vpn`)",
		Run:   runConnectionFn(),
	}

	cmd.Flags().StringP(hostFlag, "v", "", "host to use for the vpn - "+hostEnv+" at .env")
	cmd.Flags().StringP(userFlag, "u", "", "user to use - "+userEnv+" at .env")
	cmd.Flags().StringP(staticPasswordFlag, "p", "", "static password - "+staticPasswordEnv+" at .env")
	cmd.Flags().StringP(tokenKeyFlag, "t", "", "dynamic token key")
	cmd.Flags().StringP(scriptFlag, "s", "", "script to execute - "+scriptEnv+" at .env")
	cmd.Flags().StringP(openconnectAdditionalFlag, "a", "", "additional flags for openconnect - "+openconnectAdditionalEnv+" at .env")

	return cmd
}

// runConnectionFn callback to execute
func runConnectionFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString(userFlag)
		staticPassword, _ := cmd.Flags().GetString(staticPasswordFlag)
		token, _ := cmd.Flags().GetString(tokenKeyFlag)
		host, _ := cmd.Flags().GetString(hostFlag)
		script, _ := cmd.Flags().GetString(scriptFlag)
		additionalFlags, _ := cmd.Flags().GetString(openconnectAdditionalFlag)

		envAvailable := true
		err := godotenv.Load()
		if err != nil {
			envAvailable = false
			linuxUser := os.Getenv("USER")
			userSettingsFile := "/home/" + linuxUser + "/.vpn"
			err := godotenv.Load(userSettingsFile)
			if err != nil {
				fmt.Printf("\n.env neither %s file were found\n", userSettingsFile)
			} else {
				envAvailable = true
			}
		}

		if envAvailable {
			// .env data
			if host == "" {
				host = os.Getenv(hostEnv)
			}
			if user == "" {
				user = os.Getenv(userEnv)
			}
			if staticPassword == "" {
				staticPassword = os.Getenv(staticPasswordEnv)
			}
			if script == "" {
				script = os.Getenv(scriptEnv)
			}
			if additionalFlags == "" {
				additionalFlags = os.Getenv(openconnectAdditionalEnv)
			}
		}

		// read te token key if not present and connect
		if token == "" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("\nHi %s, please write your token code and press return\n", user)
			token, _ := reader.ReadString('\n')
			connect(host, user, staticPassword, token, script, additionalFlags)
		} else {
			connect(host, user, staticPassword, token, script, additionalFlags)
		}
	}
}

// Connect to the VPN using the parameters received
func connect(host string, user string, staticPassword string, token string, script string, additionalFlags string) {

	generatedCommand := "sudo openconnect " + host + " --user " + user + " -s '" + script + "' " + additionalFlags
	shCmd := "echo \"" + staticPassword + token + "\" | " + generatedCommand

	fmt.Printf("\n%s connecting to the VPN  ...  \n\n... command: %s ...\n\n", user, generatedCommand)

	cmd := exec.Command("bash", "-c", shCmd)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\n\nan error occurred, please check your information in .env | input\n\n")
		println(err.Error())
		return
	}
}
