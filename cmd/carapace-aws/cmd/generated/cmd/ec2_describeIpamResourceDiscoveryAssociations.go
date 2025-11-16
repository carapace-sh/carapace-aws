package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamResourceDiscoveryAssociationsCmd = &cobra.Command{
	Use:   "describe-ipam-resource-discovery-associations",
	Short: "Describes resource discovery association with an Amazon VPC IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamResourceDiscoveryAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIpamResourceDiscoveryAssociationsCmd).Standalone()

		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flags().String("filters", "", "The resource discovery association filters.")
		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flags().String("ipam-resource-discovery-association-ids", "", "The resource discovery association IDs.")
		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flags().String("max-results", "", "The maximum number of resource discovery associations to return in one page of results.")
		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamResourceDiscoveryAssociationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeIpamResourceDiscoveryAssociationsCmd)
}
