package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteMethodCmd = &cobra.Command{
	Use:   "delete-method",
	Short: "Deletes an existing Method resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteMethodCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteMethodCmd).Standalone()

		apigateway_deleteMethodCmd.Flags().String("http-method", "", "The HTTP verb of the Method resource.")
		apigateway_deleteMethodCmd.Flags().String("resource-id", "", "The Resource identifier for the Method resource.")
		apigateway_deleteMethodCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteMethodCmd.MarkFlagRequired("http-method")
		apigateway_deleteMethodCmd.MarkFlagRequired("resource-id")
		apigateway_deleteMethodCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteMethodCmd)
}
