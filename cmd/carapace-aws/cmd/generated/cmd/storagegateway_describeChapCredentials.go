package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeChapCredentialsCmd = &cobra.Command{
	Use:   "describe-chap-credentials",
	Short: "Returns an array of Challenge-Handshake Authentication Protocol (CHAP) credentials information for a specified iSCSI target, one for each target-initiator pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeChapCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeChapCredentialsCmd).Standalone()

		storagegateway_describeChapCredentialsCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the iSCSI volume target.")
		storagegateway_describeChapCredentialsCmd.MarkFlagRequired("target-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeChapCredentialsCmd)
}
