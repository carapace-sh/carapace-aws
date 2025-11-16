package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeScheduledInstancesCmd = &cobra.Command{
	Use:   "describe-scheduled-instances",
	Short: "Describes the specified Scheduled Instances or all your Scheduled Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeScheduledInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeScheduledInstancesCmd).Standalone()

		ec2_describeScheduledInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeScheduledInstancesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeScheduledInstancesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		ec2_describeScheduledInstancesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		ec2_describeScheduledInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeScheduledInstancesCmd.Flags().String("scheduled-instance-ids", "", "The Scheduled Instance IDs.")
		ec2_describeScheduledInstancesCmd.Flags().String("slot-start-time-range", "", "The time period for the first schedule to start.")
		ec2_describeScheduledInstancesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeScheduledInstancesCmd)
}
