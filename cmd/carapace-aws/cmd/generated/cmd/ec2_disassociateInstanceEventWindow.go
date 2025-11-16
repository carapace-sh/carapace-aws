package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateInstanceEventWindowCmd = &cobra.Command{
	Use:   "disassociate-instance-event-window",
	Short: "Disassociates one or more targets from an event window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateInstanceEventWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateInstanceEventWindowCmd).Standalone()

		ec2_disassociateInstanceEventWindowCmd.Flags().String("association-target", "", "One or more targets to disassociate from the specified event window.")
		ec2_disassociateInstanceEventWindowCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateInstanceEventWindowCmd.Flags().String("instance-event-window-id", "", "The ID of the event window.")
		ec2_disassociateInstanceEventWindowCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateInstanceEventWindowCmd.MarkFlagRequired("association-target")
		ec2_disassociateInstanceEventWindowCmd.MarkFlagRequired("instance-event-window-id")
		ec2_disassociateInstanceEventWindowCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disassociateInstanceEventWindowCmd)
}
