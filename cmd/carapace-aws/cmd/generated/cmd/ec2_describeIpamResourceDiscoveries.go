package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamResourceDiscoveriesCmd = &cobra.Command{
	Use:   "describe-ipam-resource-discoveries",
	Short: "Describes IPAM resource discoveries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamResourceDiscoveriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIpamResourceDiscoveriesCmd).Standalone()

		ec2_describeIpamResourceDiscoveriesCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamResourceDiscoveriesCmd.Flags().String("filters", "", "The resource discovery filters.")
		ec2_describeIpamResourceDiscoveriesCmd.Flags().String("ipam-resource-discovery-ids", "", "The IPAM resource discovery IDs.")
		ec2_describeIpamResourceDiscoveriesCmd.Flags().String("max-results", "", "The maximum number of resource discoveries to return in one page of results.")
		ec2_describeIpamResourceDiscoveriesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		ec2_describeIpamResourceDiscoveriesCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamResourceDiscoveriesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeIpamResourceDiscoveriesCmd)
}
