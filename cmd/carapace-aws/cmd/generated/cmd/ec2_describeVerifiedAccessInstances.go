package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVerifiedAccessInstancesCmd = &cobra.Command{
	Use:   "describe-verified-access-instances",
	Short: "Describes the specified Amazon Web Services Verified Access instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVerifiedAccessInstancesCmd).Standalone()

	ec2_describeVerifiedAccessInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVerifiedAccessInstancesCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeVerifiedAccessInstancesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeVerifiedAccessInstancesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeVerifiedAccessInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVerifiedAccessInstancesCmd.Flags().String("verified-access-instance-ids", "", "The IDs of the Verified Access instances.")
	ec2_describeVerifiedAccessInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeVerifiedAccessInstancesCmd)
}
