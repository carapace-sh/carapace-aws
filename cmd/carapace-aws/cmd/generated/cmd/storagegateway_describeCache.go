package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeCacheCmd = &cobra.Command{
	Use:   "describe-cache",
	Short: "Returns information about the cache of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeCacheCmd).Standalone()

		storagegateway_describeCacheCmd.Flags().String("gateway-arn", "", "")
		storagegateway_describeCacheCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeCacheCmd)
}
