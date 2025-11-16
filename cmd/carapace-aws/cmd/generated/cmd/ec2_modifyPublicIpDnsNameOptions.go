package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyPublicIpDnsNameOptionsCmd = &cobra.Command{
	Use:   "modify-public-ip-dns-name-options",
	Short: "Modify public hostname options for a network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyPublicIpDnsNameOptionsCmd).Standalone()

	ec2_modifyPublicIpDnsNameOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyPublicIpDnsNameOptionsCmd.Flags().String("hostname-type", "", "The public hostname type.")
	ec2_modifyPublicIpDnsNameOptionsCmd.Flags().String("network-interface-id", "", "A network interface ID.")
	ec2_modifyPublicIpDnsNameOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyPublicIpDnsNameOptionsCmd.MarkFlagRequired("hostname-type")
	ec2_modifyPublicIpDnsNameOptionsCmd.MarkFlagRequired("network-interface-id")
	ec2_modifyPublicIpDnsNameOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyPublicIpDnsNameOptionsCmd)
}
