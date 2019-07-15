package main

import (
	"github.com/alediator/vpn/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "vpn-cli"}
	rootCmd.AddCommand(cli.InitVpnCmd())
	rootCmd.Execute()
}
