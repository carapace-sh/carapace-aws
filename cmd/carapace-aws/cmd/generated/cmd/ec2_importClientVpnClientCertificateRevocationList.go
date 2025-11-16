package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_importClientVpnClientCertificateRevocationListCmd = &cobra.Command{
	Use:   "import-client-vpn-client-certificate-revocation-list",
	Short: "Uploads a client certificate revocation list to the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_importClientVpnClientCertificateRevocationListCmd).Standalone()

	ec2_importClientVpnClientCertificateRevocationListCmd.Flags().String("certificate-revocation-list", "", "The client certificate revocation list file.")
	ec2_importClientVpnClientCertificateRevocationListCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint to which the client certificate revocation list applies.")
	ec2_importClientVpnClientCertificateRevocationListCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_importClientVpnClientCertificateRevocationListCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_importClientVpnClientCertificateRevocationListCmd.MarkFlagRequired("certificate-revocation-list")
	ec2_importClientVpnClientCertificateRevocationListCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_importClientVpnClientCertificateRevocationListCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_importClientVpnClientCertificateRevocationListCmd)
}
