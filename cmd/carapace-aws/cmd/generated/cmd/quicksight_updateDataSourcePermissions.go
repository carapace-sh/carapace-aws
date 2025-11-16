package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDataSourcePermissionsCmd = &cobra.Command{
	Use:   "update-data-source-permissions",
	Short: "Updates the permissions to a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDataSourcePermissionsCmd).Standalone()

	quicksight_updateDataSourcePermissionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_updateDataSourcePermissionsCmd.Flags().String("data-source-id", "", "The ID of the data source.")
	quicksight_updateDataSourcePermissionsCmd.Flags().String("grant-permissions", "", "A list of resource permissions that you want to grant on the data source.")
	quicksight_updateDataSourcePermissionsCmd.Flags().String("revoke-permissions", "", "A list of resource permissions that you want to revoke on the data source.")
	quicksight_updateDataSourcePermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDataSourcePermissionsCmd.MarkFlagRequired("data-source-id")
	quicksightCmd.AddCommand(quicksight_updateDataSourcePermissionsCmd)
}
