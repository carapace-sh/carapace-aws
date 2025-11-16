package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getApiCmd = &cobra.Command{
	Use:   "get-api",
	Short: "Gets an Api resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getApiCmd).Standalone()

	apigatewayv2_getApiCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getApiCmd.MarkFlagRequired("api-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getApiCmd)
}
