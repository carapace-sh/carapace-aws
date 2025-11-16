package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceEventNotificationAttributesCmd = &cobra.Command{
	Use:   "describe-instance-event-notification-attributes",
	Short: "Describes the tag keys that are registered to appear in scheduled event notifications for resources in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceEventNotificationAttributesCmd).Standalone()

	ec2_describeInstanceEventNotificationAttributesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceEventNotificationAttributesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceEventNotificationAttributesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceEventNotificationAttributesCmd)
}
