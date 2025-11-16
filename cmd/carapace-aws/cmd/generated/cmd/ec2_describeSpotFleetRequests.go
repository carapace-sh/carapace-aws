package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSpotFleetRequestsCmd = &cobra.Command{
	Use:   "describe-spot-fleet-requests",
	Short: "Describes your Spot Fleet requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSpotFleetRequestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSpotFleetRequestsCmd).Standalone()

		ec2_describeSpotFleetRequestsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotFleetRequestsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSpotFleetRequestsCmd.Flags().String("next-token", "", "The token to include in another request to get the next page of items.")
		ec2_describeSpotFleetRequestsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotFleetRequestsCmd.Flags().String("spot-fleet-request-ids", "", "The IDs of the Spot Fleet requests.")
		ec2_describeSpotFleetRequestsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSpotFleetRequestsCmd)
}
