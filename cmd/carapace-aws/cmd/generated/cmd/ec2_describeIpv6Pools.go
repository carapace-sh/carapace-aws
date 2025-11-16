package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpv6PoolsCmd = &cobra.Command{
	Use:   "describe-ipv6-pools",
	Short: "Describes your IPv6 address pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpv6PoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIpv6PoolsCmd).Standalone()

		ec2_describeIpv6PoolsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeIpv6PoolsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeIpv6PoolsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeIpv6PoolsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeIpv6PoolsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeIpv6PoolsCmd.Flags().String("pool-ids", "", "The IDs of the IPv6 address pools.")
		ec2_describeIpv6PoolsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeIpv6PoolsCmd)
}
