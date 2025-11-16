package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_registerInstanceEventNotificationAttributesCmd = &cobra.Command{
	Use:   "register-instance-event-notification-attributes",
	Short: "Registers a set of tag keys to include in scheduled event notifications for your resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_registerInstanceEventNotificationAttributesCmd).Standalone()

	ec2_registerInstanceEventNotificationAttributesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_registerInstanceEventNotificationAttributesCmd.Flags().String("instance-tag-attribute", "", "Information about the tag keys to register.")
	ec2_registerInstanceEventNotificationAttributesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_registerInstanceEventNotificationAttributesCmd.MarkFlagRequired("instance-tag-attribute")
	ec2_registerInstanceEventNotificationAttributesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_registerInstanceEventNotificationAttributesCmd)
}
