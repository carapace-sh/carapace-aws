package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteApiKeyCmd = &cobra.Command{
	Use:   "delete-api-key",
	Short: "Deletes the ApiKey resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteApiKeyCmd).Standalone()

	apigateway_deleteApiKeyCmd.Flags().String("api-key", "", "The identifier of the ApiKey resource to be deleted.")
	apigateway_deleteApiKeyCmd.MarkFlagRequired("api-key")
	apigatewayCmd.AddCommand(apigateway_deleteApiKeyCmd)
}
