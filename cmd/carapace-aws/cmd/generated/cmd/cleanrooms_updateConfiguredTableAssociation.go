package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateConfiguredTableAssociationCmd = &cobra.Command{
	Use:   "update-configured-table-association",
	Short: "Updates a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateConfiguredTableAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateConfiguredTableAssociationCmd).Standalone()

		cleanrooms_updateConfiguredTableAssociationCmd.Flags().String("configured-table-association-identifier", "", "The unique identifier for the configured table association to update.")
		cleanrooms_updateConfiguredTableAssociationCmd.Flags().String("description", "", "A new description for the configured table association.")
		cleanrooms_updateConfiguredTableAssociationCmd.Flags().String("membership-identifier", "", "The unique ID for the membership that the configured table association belongs to.")
		cleanrooms_updateConfiguredTableAssociationCmd.Flags().String("role-arn", "", "The service will assume this role to access catalog metadata and query the table.")
		cleanrooms_updateConfiguredTableAssociationCmd.MarkFlagRequired("configured-table-association-identifier")
		cleanrooms_updateConfiguredTableAssociationCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateConfiguredTableAssociationCmd)
}
