package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateInstanceEventWindowCmd = &cobra.Command{
	Use:   "associate-instance-event-window",
	Short: "Associates one or more targets with an event window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateInstanceEventWindowCmd).Standalone()

	ec2_associateInstanceEventWindowCmd.Flags().String("association-target", "", "One or more targets associated with the specified event window.")
	ec2_associateInstanceEventWindowCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateInstanceEventWindowCmd.Flags().String("instance-event-window-id", "", "The ID of the event window.")
	ec2_associateInstanceEventWindowCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateInstanceEventWindowCmd.MarkFlagRequired("association-target")
	ec2_associateInstanceEventWindowCmd.MarkFlagRequired("instance-event-window-id")
	ec2_associateInstanceEventWindowCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_associateInstanceEventWindowCmd)
}
