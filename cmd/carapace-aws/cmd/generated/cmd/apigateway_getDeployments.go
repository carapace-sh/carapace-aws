package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDeploymentsCmd = &cobra.Command{
	Use:   "get-deployments",
	Short: "Gets information about a Deployments collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getDeploymentsCmd).Standalone()

		apigateway_getDeploymentsCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getDeploymentsCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getDeploymentsCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getDeploymentsCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getDeploymentsCmd)
}
