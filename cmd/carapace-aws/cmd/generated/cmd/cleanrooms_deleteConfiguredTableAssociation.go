package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteConfiguredTableAssociationCmd = &cobra.Command{
	Use:   "delete-configured-table-association",
	Short: "Deletes a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteConfiguredTableAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteConfiguredTableAssociationCmd).Standalone()

		cleanrooms_deleteConfiguredTableAssociationCmd.Flags().String("configured-table-association-identifier", "", "The unique ID for the configured table association to be deleted.")
		cleanrooms_deleteConfiguredTableAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that the configured table association belongs to.")
		cleanrooms_deleteConfiguredTableAssociationCmd.MarkFlagRequired("configured-table-association-identifier")
		cleanrooms_deleteConfiguredTableAssociationCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteConfiguredTableAssociationCmd)
}
