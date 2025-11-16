package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getDeploymentsCmd = &cobra.Command{
	Use:   "get-deployments",
	Short: "Gets the Deployments for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getDeploymentsCmd).Standalone()

		apigatewayv2_getDeploymentsCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getDeploymentsCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getDeploymentsCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getDeploymentsCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getDeploymentsCmd)
}
