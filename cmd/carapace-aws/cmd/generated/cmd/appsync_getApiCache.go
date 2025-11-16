package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getApiCacheCmd = &cobra.Command{
	Use:   "get-api-cache",
	Short: "Retrieves an `ApiCache` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getApiCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getApiCacheCmd).Standalone()

		appsync_getApiCacheCmd.Flags().String("api-id", "", "The API ID.")
		appsync_getApiCacheCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_getApiCacheCmd)
}
