package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateChapCredentialsCmd = &cobra.Command{
	Use:   "update-chap-credentials",
	Short: "Updates the Challenge-Handshake Authentication Protocol (CHAP) credentials for a specified iSCSI target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateChapCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateChapCredentialsCmd).Standalone()

		storagegateway_updateChapCredentialsCmd.Flags().String("initiator-name", "", "The iSCSI initiator that connects to the target.")
		storagegateway_updateChapCredentialsCmd.Flags().String("secret-to-authenticate-initiator", "", "The secret key that the initiator (for example, the Windows client) must provide to participate in mutual CHAP with the target.")
		storagegateway_updateChapCredentialsCmd.Flags().String("secret-to-authenticate-target", "", "The secret key that the target must provide to participate in mutual CHAP with the initiator (e.g. Windows client).")
		storagegateway_updateChapCredentialsCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the iSCSI volume target.")
		storagegateway_updateChapCredentialsCmd.MarkFlagRequired("initiator-name")
		storagegateway_updateChapCredentialsCmd.MarkFlagRequired("secret-to-authenticate-initiator")
		storagegateway_updateChapCredentialsCmd.MarkFlagRequired("target-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateChapCredentialsCmd)
}
