package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2Cmd = &cobra.Command{
	Use:   "apigatewayv2",
	Short: "Amazon API Gateway V2",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2Cmd).Standalone()

	})
	rootCmd.AddCommand(apigatewayv2Cmd)
}
