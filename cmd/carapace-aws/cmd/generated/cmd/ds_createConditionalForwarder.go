package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createConditionalForwarderCmd = &cobra.Command{
	Use:   "create-conditional-forwarder",
	Short: "Creates a conditional forwarder associated with your Amazon Web Services directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createConditionalForwarderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_createConditionalForwarderCmd).Standalone()

		ds_createConditionalForwarderCmd.Flags().String("directory-id", "", "The directory ID of the Amazon Web Services directory for which you are creating the conditional forwarder.")
		ds_createConditionalForwarderCmd.Flags().String("dns-ip-addrs", "", "The IP addresses of the remote DNS server associated with RemoteDomainName.")
		ds_createConditionalForwarderCmd.Flags().String("dns-ipv6-addrs", "", "The IPv6 addresses of the remote DNS server associated with RemoteDomainName.")
		ds_createConditionalForwarderCmd.Flags().String("remote-domain-name", "", "The fully qualified domain name (FQDN) of the remote domain with which you will set up a trust relationship.")
		ds_createConditionalForwarderCmd.MarkFlagRequired("directory-id")
		ds_createConditionalForwarderCmd.MarkFlagRequired("remote-domain-name")
	})
	dsCmd.AddCommand(ds_createConditionalForwarderCmd)
}
