package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateMethodResponseCmd = &cobra.Command{
	Use:   "update-method-response",
	Short: "Updates an existing MethodResponse resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateMethodResponseCmd).Standalone()

	apigateway_updateMethodResponseCmd.Flags().String("http-method", "", "The HTTP verb of the Method resource.")
	apigateway_updateMethodResponseCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateMethodResponseCmd.Flags().String("resource-id", "", "The Resource identifier for the MethodResponse resource.")
	apigateway_updateMethodResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateMethodResponseCmd.Flags().String("status-code", "", "The status code for the MethodResponse resource.")
	apigateway_updateMethodResponseCmd.MarkFlagRequired("http-method")
	apigateway_updateMethodResponseCmd.MarkFlagRequired("resource-id")
	apigateway_updateMethodResponseCmd.MarkFlagRequired("rest-api-id")
	apigateway_updateMethodResponseCmd.MarkFlagRequired("status-code")
	apigatewayCmd.AddCommand(apigateway_updateMethodResponseCmd)
}
