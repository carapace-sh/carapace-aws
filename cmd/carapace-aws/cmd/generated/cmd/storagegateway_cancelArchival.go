package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_cancelArchivalCmd = &cobra.Command{
	Use:   "cancel-archival",
	Short: "Cancels archiving of a virtual tape to the virtual tape shelf (VTS) after the archiving process is initiated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_cancelArchivalCmd).Standalone()

	storagegateway_cancelArchivalCmd.Flags().String("gateway-arn", "", "")
	storagegateway_cancelArchivalCmd.Flags().String("tape-arn", "", "The Amazon Resource Name (ARN) of the virtual tape you want to cancel archiving for.")
	storagegateway_cancelArchivalCmd.MarkFlagRequired("gateway-arn")
	storagegateway_cancelArchivalCmd.MarkFlagRequired("tape-arn")
	storagegatewayCmd.AddCommand(storagegateway_cancelArchivalCmd)
}
