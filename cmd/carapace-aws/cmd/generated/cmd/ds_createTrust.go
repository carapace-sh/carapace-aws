package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createTrustCmd = &cobra.Command{
	Use:   "create-trust",
	Short: "Directory Service for Microsoft Active Directory allows you to configure trust relationships.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createTrustCmd).Standalone()

	ds_createTrustCmd.Flags().String("conditional-forwarder-ip-addrs", "", "The IP addresses of the remote DNS server associated with RemoteDomainName.")
	ds_createTrustCmd.Flags().String("conditional-forwarder-ipv6-addrs", "", "The IPv6 addresses of the remote DNS server associated with RemoteDomainName.")
	ds_createTrustCmd.Flags().String("directory-id", "", "The Directory ID of the Managed Microsoft AD directory for which to establish the trust relationship.")
	ds_createTrustCmd.Flags().String("remote-domain-name", "", "The Fully Qualified Domain Name (FQDN) of the external domain for which to create the trust relationship.")
	ds_createTrustCmd.Flags().String("selective-auth", "", "Optional parameter to enable selective authentication for the trust.")
	ds_createTrustCmd.Flags().String("trust-direction", "", "The direction of the trust relationship.")
	ds_createTrustCmd.Flags().String("trust-password", "", "The trust password.")
	ds_createTrustCmd.Flags().String("trust-type", "", "The trust relationship type.")
	ds_createTrustCmd.MarkFlagRequired("directory-id")
	ds_createTrustCmd.MarkFlagRequired("remote-domain-name")
	ds_createTrustCmd.MarkFlagRequired("trust-direction")
	ds_createTrustCmd.MarkFlagRequired("trust-password")
	dsCmd.AddCommand(ds_createTrustCmd)
}
