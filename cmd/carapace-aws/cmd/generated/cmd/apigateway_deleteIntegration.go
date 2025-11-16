package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Represents a delete integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteIntegrationCmd).Standalone()

		apigateway_deleteIntegrationCmd.Flags().String("http-method", "", "Specifies a delete integration request's HTTP method.")
		apigateway_deleteIntegrationCmd.Flags().String("resource-id", "", "Specifies a delete integration request's resource identifier.")
		apigateway_deleteIntegrationCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteIntegrationCmd.MarkFlagRequired("http-method")
		apigateway_deleteIntegrationCmd.MarkFlagRequired("resource-id")
		apigateway_deleteIntegrationCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteIntegrationCmd)
}
