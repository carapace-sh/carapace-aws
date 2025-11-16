package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_putMethodResponseCmd = &cobra.Command{
	Use:   "put-method-response",
	Short: "Adds a MethodResponse to an existing Method resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_putMethodResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_putMethodResponseCmd).Standalone()

		apigateway_putMethodResponseCmd.Flags().String("http-method", "", "The HTTP verb of the Method resource.")
		apigateway_putMethodResponseCmd.Flags().String("resource-id", "", "The Resource identifier for the Method resource.")
		apigateway_putMethodResponseCmd.Flags().String("response-models", "", "Specifies the Model resources used for the response's content type.")
		apigateway_putMethodResponseCmd.Flags().String("response-parameters", "", "A key-value map specifying required or optional response parameters that API Gateway can send back to the caller.")
		apigateway_putMethodResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_putMethodResponseCmd.Flags().String("status-code", "", "The method response's status code.")
		apigateway_putMethodResponseCmd.MarkFlagRequired("http-method")
		apigateway_putMethodResponseCmd.MarkFlagRequired("resource-id")
		apigateway_putMethodResponseCmd.MarkFlagRequired("rest-api-id")
		apigateway_putMethodResponseCmd.MarkFlagRequired("status-code")
	})
	apigatewayCmd.AddCommand(apigateway_putMethodResponseCmd)
}
