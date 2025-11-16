package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateConditionalForwarderCmd = &cobra.Command{
	Use:   "update-conditional-forwarder",
	Short: "Updates a conditional forwarder that has been set up for your Amazon Web Services directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateConditionalForwarderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_updateConditionalForwarderCmd).Standalone()

		ds_updateConditionalForwarderCmd.Flags().String("directory-id", "", "The directory ID of the Amazon Web Services directory for which to update the conditional forwarder.")
		ds_updateConditionalForwarderCmd.Flags().String("dns-ip-addrs", "", "The updated IP addresses of the remote DNS server associated with the conditional forwarder.")
		ds_updateConditionalForwarderCmd.Flags().String("dns-ipv6-addrs", "", "The updated IPv6 addresses of the remote DNS server associated with the conditional forwarder.")
		ds_updateConditionalForwarderCmd.Flags().String("remote-domain-name", "", "The fully qualified domain name (FQDN) of the remote domain with which you will set up a trust relationship.")
		ds_updateConditionalForwarderCmd.MarkFlagRequired("directory-id")
		ds_updateConditionalForwarderCmd.MarkFlagRequired("remote-domain-name")
	})
	dsCmd.AddCommand(ds_updateConditionalForwarderCmd)
}
