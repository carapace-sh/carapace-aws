package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateIntegrationCmd = &cobra.Command{
	Use:   "update-integration",
	Short: "Represents an update integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateIntegrationCmd).Standalone()

	apigateway_updateIntegrationCmd.Flags().String("http-method", "", "Represents an update integration request's HTTP method.")
	apigateway_updateIntegrationCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateIntegrationCmd.Flags().String("resource-id", "", "Represents an update integration request's resource identifier.")
	apigateway_updateIntegrationCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateIntegrationCmd.MarkFlagRequired("http-method")
	apigateway_updateIntegrationCmd.MarkFlagRequired("resource-id")
	apigateway_updateIntegrationCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateIntegrationCmd)
}
