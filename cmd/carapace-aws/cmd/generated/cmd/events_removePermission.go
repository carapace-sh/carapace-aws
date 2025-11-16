package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_removePermissionCmd = &cobra.Command{
	Use:   "remove-permission",
	Short: "Revokes the permission of another Amazon Web Services account to be able to put events to the specified event bus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_removePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_removePermissionCmd).Standalone()

		events_removePermissionCmd.Flags().String("event-bus-name", "", "The name of the event bus to revoke permissions for.")
		events_removePermissionCmd.Flags().Bool("no-remove-all-permissions", false, "Specifies whether to remove all permissions.")
		events_removePermissionCmd.Flags().Bool("remove-all-permissions", false, "Specifies whether to remove all permissions.")
		events_removePermissionCmd.Flags().String("statement-id", "", "The statement ID corresponding to the account that is no longer allowed to put events to the default event bus.")
		events_removePermissionCmd.Flag("no-remove-all-permissions").Hidden = true
	})
	eventsCmd.AddCommand(events_removePermissionCmd)
}
