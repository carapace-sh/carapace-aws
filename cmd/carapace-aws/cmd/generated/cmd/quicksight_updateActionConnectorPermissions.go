package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateActionConnectorPermissionsCmd = &cobra.Command{
	Use:   "update-action-connector-permissions",
	Short: "Updates the permissions for an action connector by granting or revoking access for specific users and groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateActionConnectorPermissionsCmd).Standalone()

	quicksight_updateActionConnectorPermissionsCmd.Flags().String("action-connector-id", "", "The unique identifier of the action connector whose permissions you want to update.")
	quicksight_updateActionConnectorPermissionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID that contains the action connector.")
	quicksight_updateActionConnectorPermissionsCmd.Flags().String("grant-permissions", "", "The permissions to grant to users and groups for this action connector.")
	quicksight_updateActionConnectorPermissionsCmd.Flags().String("revoke-permissions", "", "The permissions to revoke from users and groups for this action connector.")
	quicksight_updateActionConnectorPermissionsCmd.MarkFlagRequired("action-connector-id")
	quicksight_updateActionConnectorPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_updateActionConnectorPermissionsCmd)
}
