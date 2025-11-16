package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_putPermissionCmd = &cobra.Command{
	Use:   "put-permission",
	Short: "Running `PutPermission` permits the specified Amazon Web Services account or Amazon Web Services organization to put events to the specified *event bus*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_putPermissionCmd).Standalone()

	events_putPermissionCmd.Flags().String("action", "", "The action that you are enabling the other account to perform.")
	events_putPermissionCmd.Flags().String("condition", "", "This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain Amazon Web Services organization.")
	events_putPermissionCmd.Flags().String("event-bus-name", "", "The name of the event bus associated with the rule.")
	events_putPermissionCmd.Flags().String("policy", "", "A JSON string that describes the permission policy statement.")
	events_putPermissionCmd.Flags().String("principal", "", "The 12-digit Amazon Web Services account ID that you are permitting to put events to your default event bus.")
	events_putPermissionCmd.Flags().String("statement-id", "", "An identifier string for the external account that you are granting permissions to.")
	eventsCmd.AddCommand(events_putPermissionCmd)
}
