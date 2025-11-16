package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteApiKeyCmd = &cobra.Command{
	Use:   "delete-api-key",
	Short: "Deletes an API key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteApiKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_deleteApiKeyCmd).Standalone()

		appsync_deleteApiKeyCmd.Flags().String("api-id", "", "The API ID.")
		appsync_deleteApiKeyCmd.Flags().String("id", "", "The ID for the API key.")
		appsync_deleteApiKeyCmd.MarkFlagRequired("api-id")
		appsync_deleteApiKeyCmd.MarkFlagRequired("id")
	})
	appsyncCmd.AddCommand(appsync_deleteApiKeyCmd)
}
