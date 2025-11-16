package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_retrieveTapeArchiveCmd = &cobra.Command{
	Use:   "retrieve-tape-archive",
	Short: "Retrieves an archived virtual tape from the virtual tape shelf (VTS) to a tape gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_retrieveTapeArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_retrieveTapeArchiveCmd).Standalone()

		storagegateway_retrieveTapeArchiveCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway you want to retrieve the virtual tape to.")
		storagegateway_retrieveTapeArchiveCmd.Flags().String("tape-arn", "", "The Amazon Resource Name (ARN) of the virtual tape you want to retrieve from the virtual tape shelf (VTS).")
		storagegateway_retrieveTapeArchiveCmd.MarkFlagRequired("gateway-arn")
		storagegateway_retrieveTapeArchiveCmd.MarkFlagRequired("tape-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_retrieveTapeArchiveCmd)
}
