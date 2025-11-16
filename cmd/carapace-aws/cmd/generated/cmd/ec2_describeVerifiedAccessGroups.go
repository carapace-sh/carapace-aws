package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVerifiedAccessGroupsCmd = &cobra.Command{
	Use:   "describe-verified-access-groups",
	Short: "Describes the specified Verified Access groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVerifiedAccessGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVerifiedAccessGroupsCmd).Standalone()

		ec2_describeVerifiedAccessGroupsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVerifiedAccessGroupsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeVerifiedAccessGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeVerifiedAccessGroupsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeVerifiedAccessGroupsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVerifiedAccessGroupsCmd.Flags().String("verified-access-group-ids", "", "The ID of the Verified Access groups.")
		ec2_describeVerifiedAccessGroupsCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_describeVerifiedAccessGroupsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVerifiedAccessGroupsCmd)
}
