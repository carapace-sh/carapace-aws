package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_flushApiCacheCmd = &cobra.Command{
	Use:   "flush-api-cache",
	Short: "Flushes an `ApiCache` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_flushApiCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_flushApiCacheCmd).Standalone()

		appsync_flushApiCacheCmd.Flags().String("api-id", "", "The API ID.")
		appsync_flushApiCacheCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_flushApiCacheCmd)
}
