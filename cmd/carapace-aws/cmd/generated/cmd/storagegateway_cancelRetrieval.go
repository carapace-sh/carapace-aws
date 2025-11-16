package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_cancelRetrievalCmd = &cobra.Command{
	Use:   "cancel-retrieval",
	Short: "Cancels retrieval of a virtual tape from the virtual tape shelf (VTS) to a gateway after the retrieval process is initiated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_cancelRetrievalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_cancelRetrievalCmd).Standalone()

		storagegateway_cancelRetrievalCmd.Flags().String("gateway-arn", "", "")
		storagegateway_cancelRetrievalCmd.Flags().String("tape-arn", "", "The Amazon Resource Name (ARN) of the virtual tape you want to cancel retrieval for.")
		storagegateway_cancelRetrievalCmd.MarkFlagRequired("gateway-arn")
		storagegateway_cancelRetrievalCmd.MarkFlagRequired("tape-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_cancelRetrievalCmd)
}
