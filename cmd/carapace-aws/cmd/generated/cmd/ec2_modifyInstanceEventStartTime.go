package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceEventStartTimeCmd = &cobra.Command{
	Use:   "modify-instance-event-start-time",
	Short: "Modifies the start time for a scheduled Amazon EC2 instance event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceEventStartTimeCmd).Standalone()

	ec2_modifyInstanceEventStartTimeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceEventStartTimeCmd.Flags().String("instance-event-id", "", "The ID of the event whose date and time you are modifying.")
	ec2_modifyInstanceEventStartTimeCmd.Flags().String("instance-id", "", "The ID of the instance with the scheduled event.")
	ec2_modifyInstanceEventStartTimeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceEventStartTimeCmd.Flags().String("not-before", "", "The new date and time when the event will take place.")
	ec2_modifyInstanceEventStartTimeCmd.MarkFlagRequired("instance-event-id")
	ec2_modifyInstanceEventStartTimeCmd.MarkFlagRequired("instance-id")
	ec2_modifyInstanceEventStartTimeCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyInstanceEventStartTimeCmd.MarkFlagRequired("not-before")
	ec2Cmd.AddCommand(ec2_modifyInstanceEventStartTimeCmd)
}
