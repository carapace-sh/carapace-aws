package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateApiKeyCmd = &cobra.Command{
	Use:   "update-api-key",
	Short: "Updates an API key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateApiKeyCmd).Standalone()

	appsync_updateApiKeyCmd.Flags().String("api-id", "", "The ID for the GraphQL API.")
	appsync_updateApiKeyCmd.Flags().String("description", "", "A description of the purpose of the API key.")
	appsync_updateApiKeyCmd.Flags().String("expires", "", "From the update time, the time after which the API key expires.")
	appsync_updateApiKeyCmd.Flags().String("id", "", "The API key ID.")
	appsync_updateApiKeyCmd.MarkFlagRequired("api-id")
	appsync_updateApiKeyCmd.MarkFlagRequired("id")
	appsyncCmd.AddCommand(appsync_updateApiKeyCmd)
}
