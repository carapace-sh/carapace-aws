package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeListenerCertificatesCmd = &cobra.Command{
	Use:   "describe-listener-certificates",
	Short: "Describes the default certificate and the certificate list for the specified HTTPS or TLS listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeListenerCertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeListenerCertificatesCmd).Standalone()

		elbv2_describeListenerCertificatesCmd.Flags().String("listener-arn", "", "The Amazon Resource Names (ARN) of the listener.")
		elbv2_describeListenerCertificatesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		elbv2_describeListenerCertificatesCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
		elbv2_describeListenerCertificatesCmd.MarkFlagRequired("listener-arn")
	})
	elbv2Cmd.AddCommand(elbv2_describeListenerCertificatesCmd)
}
