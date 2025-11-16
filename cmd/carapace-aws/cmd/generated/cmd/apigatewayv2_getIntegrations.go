package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getIntegrationsCmd = &cobra.Command{
	Use:   "get-integrations",
	Short: "Gets the Integrations for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getIntegrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getIntegrationsCmd).Standalone()

		apigatewayv2_getIntegrationsCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getIntegrationsCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getIntegrationsCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getIntegrationsCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getIntegrationsCmd)
}
