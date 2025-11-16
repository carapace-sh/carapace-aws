package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteMethodResponseCmd = &cobra.Command{
	Use:   "delete-method-response",
	Short: "Deletes an existing MethodResponse resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteMethodResponseCmd).Standalone()

	apigateway_deleteMethodResponseCmd.Flags().String("http-method", "", "The HTTP verb of the Method resource.")
	apigateway_deleteMethodResponseCmd.Flags().String("resource-id", "", "The Resource identifier for the MethodResponse resource.")
	apigateway_deleteMethodResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_deleteMethodResponseCmd.Flags().String("status-code", "", "The status code identifier for the MethodResponse resource.")
	apigateway_deleteMethodResponseCmd.MarkFlagRequired("http-method")
	apigateway_deleteMethodResponseCmd.MarkFlagRequired("resource-id")
	apigateway_deleteMethodResponseCmd.MarkFlagRequired("rest-api-id")
	apigateway_deleteMethodResponseCmd.MarkFlagRequired("status-code")
	apigatewayCmd.AddCommand(apigateway_deleteMethodResponseCmd)
}
