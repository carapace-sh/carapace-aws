package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createApiKeyCmd = &cobra.Command{
	Use:   "create-api-key",
	Short: "Creates a unique key that you can distribute to clients who invoke your API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createApiKeyCmd).Standalone()

	appsync_createApiKeyCmd.Flags().String("api-id", "", "The ID for your GraphQL API.")
	appsync_createApiKeyCmd.Flags().String("description", "", "A description of the purpose of the API key.")
	appsync_createApiKeyCmd.Flags().String("expires", "", "From the creation time, the time after which the API key expires.")
	appsync_createApiKeyCmd.MarkFlagRequired("api-id")
	appsyncCmd.AddCommand(appsync_createApiKeyCmd)
}
