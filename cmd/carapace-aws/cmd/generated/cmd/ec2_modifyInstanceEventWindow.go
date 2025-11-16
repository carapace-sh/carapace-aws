package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceEventWindowCmd = &cobra.Command{
	Use:   "modify-instance-event-window",
	Short: "Modifies the specified event window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceEventWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyInstanceEventWindowCmd).Standalone()

		ec2_modifyInstanceEventWindowCmd.Flags().String("cron-expression", "", "The cron expression of the event window, for example, `* 0-4,20-23 * * 1,5`.")
		ec2_modifyInstanceEventWindowCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyInstanceEventWindowCmd.Flags().String("instance-event-window-id", "", "The ID of the event window.")
		ec2_modifyInstanceEventWindowCmd.Flags().String("name", "", "The name of the event window.")
		ec2_modifyInstanceEventWindowCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyInstanceEventWindowCmd.Flags().String("time-ranges", "", "The time ranges of the event window.")
		ec2_modifyInstanceEventWindowCmd.MarkFlagRequired("instance-event-window-id")
		ec2_modifyInstanceEventWindowCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyInstanceEventWindowCmd)
}
