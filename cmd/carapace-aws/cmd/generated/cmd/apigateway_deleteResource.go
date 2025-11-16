package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteResourceCmd = &cobra.Command{
	Use:   "delete-resource",
	Short: "Deletes a Resource resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteResourceCmd).Standalone()

		apigateway_deleteResourceCmd.Flags().String("resource-id", "", "The identifier of the Resource resource.")
		apigateway_deleteResourceCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteResourceCmd.MarkFlagRequired("resource-id")
		apigateway_deleteResourceCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteResourceCmd)
}
