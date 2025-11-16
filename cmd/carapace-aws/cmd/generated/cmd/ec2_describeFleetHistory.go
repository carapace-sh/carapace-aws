package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFleetHistoryCmd = &cobra.Command{
	Use:   "describe-fleet-history",
	Short: "Describes the events for the specified EC2 Fleet during the specified time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFleetHistoryCmd).Standalone()

	ec2_describeFleetHistoryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFleetHistoryCmd.Flags().String("event-type", "", "The type of events to describe.")
	ec2_describeFleetHistoryCmd.Flags().String("fleet-id", "", "The ID of the EC2 Fleet.")
	ec2_describeFleetHistoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeFleetHistoryCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeFleetHistoryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFleetHistoryCmd.Flags().String("start-time", "", "The start date and time for the events, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
	ec2_describeFleetHistoryCmd.MarkFlagRequired("fleet-id")
	ec2_describeFleetHistoryCmd.Flag("no-dry-run").Hidden = true
	ec2_describeFleetHistoryCmd.MarkFlagRequired("start-time")
	ec2Cmd.AddCommand(ec2_describeFleetHistoryCmd)
}
