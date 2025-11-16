package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeScheduledInstanceAvailabilityCmd = &cobra.Command{
	Use:   "describe-scheduled-instance-availability",
	Short: "Finds available schedules that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeScheduledInstanceAvailabilityCmd).Standalone()

	ec2_describeScheduledInstanceAvailabilityCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("filters", "", "The filters.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("first-slot-start-time-range", "", "The time period for the first schedule to start.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("max-slot-duration-in-hours", "", "The maximum available duration, in hours.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("min-slot-duration-in-hours", "", "The minimum available duration, in hours.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("next-token", "", "The token for the next set of results.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeScheduledInstanceAvailabilityCmd.Flags().String("recurrence", "", "The schedule recurrence.")
	ec2_describeScheduledInstanceAvailabilityCmd.MarkFlagRequired("first-slot-start-time-range")
	ec2_describeScheduledInstanceAvailabilityCmd.Flag("no-dry-run").Hidden = true
	ec2_describeScheduledInstanceAvailabilityCmd.MarkFlagRequired("recurrence")
	ec2Cmd.AddCommand(ec2_describeScheduledInstanceAvailabilityCmd)
}
