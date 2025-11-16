package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_disassociateOrganizationalUnitCmd = &cobra.Command{
	Use:   "disassociate-organizational-unit",
	Short: "Removes the association between an organizational unit and a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_disassociateOrganizationalUnitCmd).Standalone()

	notifications_disassociateOrganizationalUnitCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the notification configuration to disassociate from the organizational unit.")
	notifications_disassociateOrganizationalUnitCmd.Flags().String("organizational-unit-id", "", "The unique identifier of the organizational unit to disassociate.")
	notifications_disassociateOrganizationalUnitCmd.MarkFlagRequired("notification-configuration-arn")
	notifications_disassociateOrganizationalUnitCmd.MarkFlagRequired("organizational-unit-id")
	notificationsCmd.AddCommand(notifications_disassociateOrganizationalUnitCmd)
}
