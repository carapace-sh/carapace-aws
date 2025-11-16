package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_disableNotificationsAccessForOrganizationCmd = &cobra.Command{
	Use:   "disable-notifications-access-for-organization",
	Short: "Disables service trust between User Notifications and Amazon Web Services Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_disableNotificationsAccessForOrganizationCmd).Standalone()

	notificationsCmd.AddCommand(notifications_disableNotificationsAccessForOrganizationCmd)
}
