package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getConfiguredTableAssociationCmd = &cobra.Command{
	Use:   "get-configured-table-association",
	Short: "Retrieves a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getConfiguredTableAssociationCmd).Standalone()

	cleanrooms_getConfiguredTableAssociationCmd.Flags().String("configured-table-association-identifier", "", "The unique ID for the configured table association to retrieve.")
	cleanrooms_getConfiguredTableAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that the configured table association belongs to.")
	cleanrooms_getConfiguredTableAssociationCmd.MarkFlagRequired("configured-table-association-identifier")
	cleanrooms_getConfiguredTableAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getConfiguredTableAssociationCmd)
}
