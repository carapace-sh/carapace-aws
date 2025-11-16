package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_removeListenerCertificatesCmd = &cobra.Command{
	Use:   "remove-listener-certificates",
	Short: "Removes the specified certificate from the certificate list for the specified HTTPS or TLS listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_removeListenerCertificatesCmd).Standalone()

	elbv2_removeListenerCertificatesCmd.Flags().String("certificates", "", "The certificate to remove.")
	elbv2_removeListenerCertificatesCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
	elbv2_removeListenerCertificatesCmd.MarkFlagRequired("certificates")
	elbv2_removeListenerCertificatesCmd.MarkFlagRequired("listener-arn")
	elbv2Cmd.AddCommand(elbv2_removeListenerCertificatesCmd)
}
