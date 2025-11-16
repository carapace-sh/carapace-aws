package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamPoolsCmd = &cobra.Command{
	Use:   "describe-ipam-pools",
	Short: "Get information about your IPAM pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIpamPoolsCmd).Standalone()

		ec2_describeIpamPoolsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamPoolsCmd.Flags().String("filters", "", "One or more filters for the request.")
		ec2_describeIpamPoolsCmd.Flags().String("ipam-pool-ids", "", "The IDs of the IPAM pools you would like information on.")
		ec2_describeIpamPoolsCmd.Flags().String("max-results", "", "The maximum number of results to return in the request.")
		ec2_describeIpamPoolsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeIpamPoolsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamPoolsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeIpamPoolsCmd)
}
