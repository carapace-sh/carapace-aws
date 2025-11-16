package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteIdMappingTableCmd = &cobra.Command{
	Use:   "delete-id-mapping-table",
	Short: "Deletes an ID mapping table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteIdMappingTableCmd).Standalone()

	cleanrooms_deleteIdMappingTableCmd.Flags().String("id-mapping-table-identifier", "", "The unique identifier of the ID mapping table that you want to delete.")
	cleanrooms_deleteIdMappingTableCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID mapping table that you want to delete.")
	cleanrooms_deleteIdMappingTableCmd.MarkFlagRequired("id-mapping-table-identifier")
	cleanrooms_deleteIdMappingTableCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_deleteIdMappingTableCmd)
}
