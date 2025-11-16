package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_addCacheCmd = &cobra.Command{
	Use:   "add-cache",
	Short: "Configures one or more gateway local disks as cache for a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_addCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_addCacheCmd).Standalone()

		storagegateway_addCacheCmd.Flags().String("disk-ids", "", "An array of strings that identify disks that are to be configured as working storage.")
		storagegateway_addCacheCmd.Flags().String("gateway-arn", "", "")
		storagegateway_addCacheCmd.MarkFlagRequired("disk-ids")
		storagegateway_addCacheCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_addCacheCmd)
}
