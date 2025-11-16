package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_associateOrganizationalUnitCmd = &cobra.Command{
	Use:   "associate-organizational-unit",
	Short: "Associates an organizational unit with a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_associateOrganizationalUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_associateOrganizationalUnitCmd).Standalone()

		notifications_associateOrganizationalUnitCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the notification configuration to associate with the organizational unit.")
		notifications_associateOrganizationalUnitCmd.Flags().String("organizational-unit-id", "", "The unique identifier of the organizational unit to associate.")
		notifications_associateOrganizationalUnitCmd.MarkFlagRequired("notification-configuration-arn")
		notifications_associateOrganizationalUnitCmd.MarkFlagRequired("organizational-unit-id")
	})
	notificationsCmd.AddCommand(notifications_associateOrganizationalUnitCmd)
}
