package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSpotFleetInstancesCmd = &cobra.Command{
	Use:   "describe-spot-fleet-instances",
	Short: "Describes the running instances for the specified Spot Fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSpotFleetInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSpotFleetInstancesCmd).Standalone()

		ec2_describeSpotFleetInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotFleetInstancesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSpotFleetInstancesCmd.Flags().String("next-token", "", "The token to include in another request to get the next page of items.")
		ec2_describeSpotFleetInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotFleetInstancesCmd.Flags().String("spot-fleet-request-id", "", "The ID of the Spot Fleet request.")
		ec2_describeSpotFleetInstancesCmd.Flag("no-dry-run").Hidden = true
		ec2_describeSpotFleetInstancesCmd.MarkFlagRequired("spot-fleet-request-id")
	})
	ec2Cmd.AddCommand(ec2_describeSpotFleetInstancesCmd)
}
