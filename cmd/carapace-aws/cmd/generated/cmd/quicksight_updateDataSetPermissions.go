package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDataSetPermissionsCmd = &cobra.Command{
	Use:   "update-data-set-permissions",
	Short: "Updates the permissions on a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDataSetPermissionsCmd).Standalone()

	quicksight_updateDataSetPermissionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_updateDataSetPermissionsCmd.Flags().String("data-set-id", "", "The ID for the dataset whose permissions you want to update.")
	quicksight_updateDataSetPermissionsCmd.Flags().String("grant-permissions", "", "The resource permissions that you want to grant to the dataset.")
	quicksight_updateDataSetPermissionsCmd.Flags().String("revoke-permissions", "", "The resource permissions that you want to revoke from the dataset.")
	quicksight_updateDataSetPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDataSetPermissionsCmd.MarkFlagRequired("data-set-id")
	quicksightCmd.AddCommand(quicksight_updateDataSetPermissionsCmd)
}
