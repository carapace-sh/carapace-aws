package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_resetCacheCmd = &cobra.Command{
	Use:   "reset-cache",
	Short: "Resets all cache disks that have encountered an error and makes the disks available for reconfiguration as cache storage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_resetCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_resetCacheCmd).Standalone()

		storagegateway_resetCacheCmd.Flags().String("gateway-arn", "", "")
		storagegateway_resetCacheCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_resetCacheCmd)
}
