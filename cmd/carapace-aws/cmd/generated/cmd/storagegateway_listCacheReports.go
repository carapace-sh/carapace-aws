package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listCacheReportsCmd = &cobra.Command{
	Use:   "list-cache-reports",
	Short: "Returns a list of existing cache reports for all file shares associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listCacheReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listCacheReportsCmd).Standalone()

		storagegateway_listCacheReportsCmd.Flags().String("marker", "", "Opaque pagination token returned from a previous `ListCacheReports` operation.")
	})
	storagegatewayCmd.AddCommand(storagegateway_listCacheReportsCmd)
}
