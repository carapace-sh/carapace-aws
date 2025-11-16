package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deregisterInstanceEventNotificationAttributesCmd = &cobra.Command{
	Use:   "deregister-instance-event-notification-attributes",
	Short: "Deregisters tag keys to prevent tags that have the specified tag keys from being included in scheduled event notifications for resources in the Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deregisterInstanceEventNotificationAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deregisterInstanceEventNotificationAttributesCmd).Standalone()

		ec2_deregisterInstanceEventNotificationAttributesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterInstanceEventNotificationAttributesCmd.Flags().String("instance-tag-attribute", "", "Information about the tag keys to deregister.")
		ec2_deregisterInstanceEventNotificationAttributesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterInstanceEventNotificationAttributesCmd.MarkFlagRequired("instance-tag-attribute")
		ec2_deregisterInstanceEventNotificationAttributesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deregisterInstanceEventNotificationAttributesCmd)
}
