package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_runScheduledInstancesCmd = &cobra.Command{
	Use:   "run-scheduled-instances",
	Short: "Launches the specified Scheduled Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_runScheduledInstancesCmd).Standalone()

	ec2_runScheduledInstancesCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that ensures the idempotency of the request.")
	ec2_runScheduledInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_runScheduledInstancesCmd.Flags().String("instance-count", "", "The number of instances.")
	ec2_runScheduledInstancesCmd.Flags().String("launch-specification", "", "The launch specification.")
	ec2_runScheduledInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_runScheduledInstancesCmd.Flags().String("scheduled-instance-id", "", "The Scheduled Instance ID.")
	ec2_runScheduledInstancesCmd.MarkFlagRequired("launch-specification")
	ec2_runScheduledInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2_runScheduledInstancesCmd.MarkFlagRequired("scheduled-instance-id")
	ec2Cmd.AddCommand(ec2_runScheduledInstancesCmd)
}
