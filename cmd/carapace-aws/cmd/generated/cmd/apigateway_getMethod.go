package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getMethodCmd = &cobra.Command{
	Use:   "get-method",
	Short: "Describe an existing Method resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getMethodCmd).Standalone()

	apigateway_getMethodCmd.Flags().String("http-method", "", "Specifies the method request's HTTP method type.")
	apigateway_getMethodCmd.Flags().String("resource-id", "", "The Resource identifier for the Method resource.")
	apigateway_getMethodCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getMethodCmd.MarkFlagRequired("http-method")
	apigateway_getMethodCmd.MarkFlagRequired("resource-id")
	apigateway_getMethodCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getMethodCmd)
}
