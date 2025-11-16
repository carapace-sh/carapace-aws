package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVerifiedAccessEndpointsCmd = &cobra.Command{
	Use:   "describe-verified-access-endpoints",
	Short: "Describes the specified Amazon Web Services Verified Access endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVerifiedAccessEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVerifiedAccessEndpointsCmd).Standalone()

		ec2_describeVerifiedAccessEndpointsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().String("verified-access-endpoint-ids", "", "The ID of the Verified Access endpoint.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group.")
		ec2_describeVerifiedAccessEndpointsCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_describeVerifiedAccessEndpointsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVerifiedAccessEndpointsCmd)
}
