package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSpotFleetRequestHistoryCmd = &cobra.Command{
	Use:   "describe-spot-fleet-request-history",
	Short: "Describes the events for the specified Spot Fleet request during the specified time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSpotFleetRequestHistoryCmd).Standalone()

	ec2_describeSpotFleetRequestHistoryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSpotFleetRequestHistoryCmd.Flags().String("event-type", "", "The type of events to describe.")
	ec2_describeSpotFleetRequestHistoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeSpotFleetRequestHistoryCmd.Flags().String("next-token", "", "The token to include in another request to get the next page of items.")
	ec2_describeSpotFleetRequestHistoryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSpotFleetRequestHistoryCmd.Flags().String("spot-fleet-request-id", "", "The ID of the Spot Fleet request.")
	ec2_describeSpotFleetRequestHistoryCmd.Flags().String("start-time", "", "The starting date and time for the events, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
	ec2_describeSpotFleetRequestHistoryCmd.Flag("no-dry-run").Hidden = true
	ec2_describeSpotFleetRequestHistoryCmd.MarkFlagRequired("spot-fleet-request-id")
	ec2_describeSpotFleetRequestHistoryCmd.MarkFlagRequired("start-time")
	ec2Cmd.AddCommand(ec2_describeSpotFleetRequestHistoryCmd)
}
