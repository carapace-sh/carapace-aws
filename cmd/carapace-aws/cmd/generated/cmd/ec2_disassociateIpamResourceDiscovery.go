package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateIpamResourceDiscoveryCmd = &cobra.Command{
	Use:   "disassociate-ipam-resource-discovery",
	Short: "Disassociates a resource discovery from an Amazon VPC IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateIpamResourceDiscoveryCmd).Standalone()

	ec2_disassociateIpamResourceDiscoveryCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_disassociateIpamResourceDiscoveryCmd.Flags().String("ipam-resource-discovery-association-id", "", "A resource discovery association ID.")
	ec2_disassociateIpamResourceDiscoveryCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_disassociateIpamResourceDiscoveryCmd.MarkFlagRequired("ipam-resource-discovery-association-id")
	ec2_disassociateIpamResourceDiscoveryCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disassociateIpamResourceDiscoveryCmd)
}
