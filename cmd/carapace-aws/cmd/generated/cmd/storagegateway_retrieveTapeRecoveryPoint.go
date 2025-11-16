package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_retrieveTapeRecoveryPointCmd = &cobra.Command{
	Use:   "retrieve-tape-recovery-point",
	Short: "Retrieves the recovery point for the specified virtual tape.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_retrieveTapeRecoveryPointCmd).Standalone()

	storagegateway_retrieveTapeRecoveryPointCmd.Flags().String("gateway-arn", "", "")
	storagegateway_retrieveTapeRecoveryPointCmd.Flags().String("tape-arn", "", "The Amazon Resource Name (ARN) of the virtual tape for which you want to retrieve the recovery point.")
	storagegateway_retrieveTapeRecoveryPointCmd.MarkFlagRequired("gateway-arn")
	storagegateway_retrieveTapeRecoveryPointCmd.MarkFlagRequired("tape-arn")
	storagegatewayCmd.AddCommand(storagegateway_retrieveTapeRecoveryPointCmd)
}
