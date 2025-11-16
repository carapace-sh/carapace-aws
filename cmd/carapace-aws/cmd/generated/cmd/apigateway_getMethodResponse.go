package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getMethodResponseCmd = &cobra.Command{
	Use:   "get-method-response",
	Short: "Describes a MethodResponse resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getMethodResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getMethodResponseCmd).Standalone()

		apigateway_getMethodResponseCmd.Flags().String("http-method", "", "The HTTP verb of the Method resource.")
		apigateway_getMethodResponseCmd.Flags().String("resource-id", "", "The Resource identifier for the MethodResponse resource.")
		apigateway_getMethodResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getMethodResponseCmd.Flags().String("status-code", "", "The status code for the MethodResponse resource.")
		apigateway_getMethodResponseCmd.MarkFlagRequired("http-method")
		apigateway_getMethodResponseCmd.MarkFlagRequired("resource-id")
		apigateway_getMethodResponseCmd.MarkFlagRequired("rest-api-id")
		apigateway_getMethodResponseCmd.MarkFlagRequired("status-code")
	})
	apigatewayCmd.AddCommand(apigateway_getMethodResponseCmd)
}
