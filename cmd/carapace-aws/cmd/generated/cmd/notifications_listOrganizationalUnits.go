package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listOrganizationalUnitsCmd = &cobra.Command{
	Use:   "list-organizational-units",
	Short: "Returns a list of organizational units associated with a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listOrganizationalUnitsCmd).Standalone()

	notifications_listOrganizationalUnitsCmd.Flags().String("max-results", "", "The maximum number of organizational units to return in a single call.")
	notifications_listOrganizationalUnitsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	notifications_listOrganizationalUnitsCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the notification configuration used to filter the organizational units.")
	notifications_listOrganizationalUnitsCmd.MarkFlagRequired("notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_listOrganizationalUnitsCmd)
}
