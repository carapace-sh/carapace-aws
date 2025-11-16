package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getAccountCmd = &cobra.Command{
	Use:   "get-account",
	Short: "Gets information about the current Account resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getAccountCmd).Standalone()

	apigatewayCmd.AddCommand(apigateway_getAccountCmd)
}
