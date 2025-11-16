package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createInstanceEventWindowCmd = &cobra.Command{
	Use:   "create-instance-event-window",
	Short: "Creates an event window in which scheduled events for the associated Amazon EC2 instances can run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createInstanceEventWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createInstanceEventWindowCmd).Standalone()

		ec2_createInstanceEventWindowCmd.Flags().String("cron-expression", "", "The cron expression for the event window, for example, `* 0-4,20-23 * * 1,5`.")
		ec2_createInstanceEventWindowCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createInstanceEventWindowCmd.Flags().String("name", "", "The name of the event window.")
		ec2_createInstanceEventWindowCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createInstanceEventWindowCmd.Flags().String("tag-specifications", "", "The tags to apply to the event window.")
		ec2_createInstanceEventWindowCmd.Flags().String("time-ranges", "", "The time range for the event window.")
		ec2_createInstanceEventWindowCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createInstanceEventWindowCmd)
}
