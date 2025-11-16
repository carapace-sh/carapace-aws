package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getIdMappingTableCmd = &cobra.Command{
	Use:   "get-id-mapping-table",
	Short: "Retrieves an ID mapping table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getIdMappingTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getIdMappingTableCmd).Standalone()

		cleanrooms_getIdMappingTableCmd.Flags().String("id-mapping-table-identifier", "", "The unique identifier of the ID mapping table identifier that you want to retrieve.")
		cleanrooms_getIdMappingTableCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID mapping table that you want to retrieve.")
		cleanrooms_getIdMappingTableCmd.MarkFlagRequired("id-mapping-table-identifier")
		cleanrooms_getIdMappingTableCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getIdMappingTableCmd)
}
