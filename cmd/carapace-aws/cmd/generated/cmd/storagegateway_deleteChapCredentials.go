package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteChapCredentialsCmd = &cobra.Command{
	Use:   "delete-chap-credentials",
	Short: "Deletes Challenge-Handshake Authentication Protocol (CHAP) credentials for a specified iSCSI target and initiator pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteChapCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteChapCredentialsCmd).Standalone()

		storagegateway_deleteChapCredentialsCmd.Flags().String("initiator-name", "", "The iSCSI initiator that connects to the target.")
		storagegateway_deleteChapCredentialsCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the iSCSI volume target.")
		storagegateway_deleteChapCredentialsCmd.MarkFlagRequired("initiator-name")
		storagegateway_deleteChapCredentialsCmd.MarkFlagRequired("target-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteChapCredentialsCmd)
}
