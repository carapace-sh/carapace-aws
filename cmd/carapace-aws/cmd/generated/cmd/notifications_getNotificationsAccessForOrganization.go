package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getNotificationsAccessForOrganizationCmd = &cobra.Command{
	Use:   "get-notifications-access-for-organization",
	Short: "Returns the AccessStatus of Service Trust Enablement for User Notifications and Amazon Web Services Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getNotificationsAccessForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_getNotificationsAccessForOrganizationCmd).Standalone()

	})
	notificationsCmd.AddCommand(notifications_getNotificationsAccessForOrganizationCmd)
}
