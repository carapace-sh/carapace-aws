package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_addListenerCertificatesCmd = &cobra.Command{
	Use:   "add-listener-certificates",
	Short: "Adds the specified SSL server certificate to the certificate list for the specified HTTPS or TLS listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_addListenerCertificatesCmd).Standalone()

	elbv2_addListenerCertificatesCmd.Flags().String("certificates", "", "The certificate to add.")
	elbv2_addListenerCertificatesCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
	elbv2_addListenerCertificatesCmd.MarkFlagRequired("certificates")
	elbv2_addListenerCertificatesCmd.MarkFlagRequired("listener-arn")
	elbv2Cmd.AddCommand(elbv2_addListenerCertificatesCmd)
}
