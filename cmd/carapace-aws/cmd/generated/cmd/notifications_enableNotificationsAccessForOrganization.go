package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_enableNotificationsAccessForOrganizationCmd = &cobra.Command{
	Use:   "enable-notifications-access-for-organization",
	Short: "Enables service trust between User Notifications and Amazon Web Services Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_enableNotificationsAccessForOrganizationCmd).Standalone()

	notificationsCmd.AddCommand(notifications_enableNotificationsAccessForOrganizationCmd)
}
