package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteInstanceEventWindowCmd = &cobra.Command{
	Use:   "delete-instance-event-window",
	Short: "Deletes the specified event window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteInstanceEventWindowCmd).Standalone()

	ec2_deleteInstanceEventWindowCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteInstanceEventWindowCmd.Flags().Bool("force-delete", false, "Specify `true` to force delete the event window.")
	ec2_deleteInstanceEventWindowCmd.Flags().String("instance-event-window-id", "", "The ID of the event window.")
	ec2_deleteInstanceEventWindowCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteInstanceEventWindowCmd.Flags().Bool("no-force-delete", false, "Specify `true` to force delete the event window.")
	ec2_deleteInstanceEventWindowCmd.MarkFlagRequired("instance-event-window-id")
	ec2_deleteInstanceEventWindowCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteInstanceEventWindowCmd.Flag("no-force-delete").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteInstanceEventWindowCmd)
}
