package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_refreshCacheCmd = &cobra.Command{
	Use:   "refresh-cache",
	Short: "Refreshes the cached inventory of objects for the specified file share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_refreshCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_refreshCacheCmd).Standalone()

		storagegateway_refreshCacheCmd.Flags().String("file-share-arn", "", "The Amazon Resource Name (ARN) of the file share you want to refresh.")
		storagegateway_refreshCacheCmd.Flags().String("folder-list", "", "A comma-separated list of the paths of folders to refresh in the cache.")
		storagegateway_refreshCacheCmd.Flags().Bool("no-recursive", false, "A value that specifies whether to recursively refresh folders in the cache.")
		storagegateway_refreshCacheCmd.Flags().Bool("recursive", false, "A value that specifies whether to recursively refresh folders in the cache.")
		storagegateway_refreshCacheCmd.MarkFlagRequired("file-share-arn")
		storagegateway_refreshCacheCmd.Flag("no-recursive").Hidden = true
	})
	storagegatewayCmd.AddCommand(storagegateway_refreshCacheCmd)
}
