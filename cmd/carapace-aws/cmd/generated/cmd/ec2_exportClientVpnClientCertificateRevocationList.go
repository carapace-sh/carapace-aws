package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_exportClientVpnClientCertificateRevocationListCmd = &cobra.Command{
	Use:   "export-client-vpn-client-certificate-revocation-list",
	Short: "Downloads the client certificate revocation list for the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_exportClientVpnClientCertificateRevocationListCmd).Standalone()

	ec2_exportClientVpnClientCertificateRevocationListCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_exportClientVpnClientCertificateRevocationListCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportClientVpnClientCertificateRevocationListCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportClientVpnClientCertificateRevocationListCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_exportClientVpnClientCertificateRevocationListCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_exportClientVpnClientCertificateRevocationListCmd)
}
