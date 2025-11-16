package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCoipPoolsCmd = &cobra.Command{
	Use:   "describe-coip-pools",
	Short: "Describes the specified customer-owned address pools or all of your customer-owned address pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCoipPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCoipPoolsCmd).Standalone()

		ec2_describeCoipPoolsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCoipPoolsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeCoipPoolsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeCoipPoolsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeCoipPoolsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCoipPoolsCmd.Flags().String("pool-ids", "", "The IDs of the address pools.")
		ec2_describeCoipPoolsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCoipPoolsCmd)
}
