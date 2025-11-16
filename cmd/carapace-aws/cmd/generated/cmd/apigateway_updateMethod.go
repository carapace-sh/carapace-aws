package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateMethodCmd = &cobra.Command{
	Use:   "update-method",
	Short: "Updates an existing Method resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateMethodCmd).Standalone()

	apigateway_updateMethodCmd.Flags().String("http-method", "", "The HTTP verb of the Method resource.")
	apigateway_updateMethodCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateMethodCmd.Flags().String("resource-id", "", "The Resource identifier for the Method resource.")
	apigateway_updateMethodCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateMethodCmd.MarkFlagRequired("http-method")
	apigateway_updateMethodCmd.MarkFlagRequired("resource-id")
	apigateway_updateMethodCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateMethodCmd)
}
