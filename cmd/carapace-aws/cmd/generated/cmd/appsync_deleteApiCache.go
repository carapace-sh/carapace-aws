package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteApiCacheCmd = &cobra.Command{
	Use:   "delete-api-cache",
	Short: "Deletes an `ApiCache` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteApiCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_deleteApiCacheCmd).Standalone()

		appsync_deleteApiCacheCmd.Flags().String("api-id", "", "The API ID.")
		appsync_deleteApiCacheCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_deleteApiCacheCmd)
}
