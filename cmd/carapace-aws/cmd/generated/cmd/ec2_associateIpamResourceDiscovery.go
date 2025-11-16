package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateIpamResourceDiscoveryCmd = &cobra.Command{
	Use:   "associate-ipam-resource-discovery",
	Short: "Associates an IPAM resource discovery with an Amazon VPC IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateIpamResourceDiscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateIpamResourceDiscoveryCmd).Standalone()

		ec2_associateIpamResourceDiscoveryCmd.Flags().String("client-token", "", "A client token.")
		ec2_associateIpamResourceDiscoveryCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_associateIpamResourceDiscoveryCmd.Flags().String("ipam-id", "", "An IPAM ID.")
		ec2_associateIpamResourceDiscoveryCmd.Flags().String("ipam-resource-discovery-id", "", "A resource discovery ID.")
		ec2_associateIpamResourceDiscoveryCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_associateIpamResourceDiscoveryCmd.Flags().String("tag-specifications", "", "Tag specifications.")
		ec2_associateIpamResourceDiscoveryCmd.MarkFlagRequired("ipam-id")
		ec2_associateIpamResourceDiscoveryCmd.MarkFlagRequired("ipam-resource-discovery-id")
		ec2_associateIpamResourceDiscoveryCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_associateIpamResourceDiscoveryCmd)
}
