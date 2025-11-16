package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listIdMappingTablesCmd = &cobra.Command{
	Use:   "list-id-mapping-tables",
	Short: "Returns a list of ID mapping tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listIdMappingTablesCmd).Standalone()

	cleanrooms_listIdMappingTablesCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanrooms_listIdMappingTablesCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID mapping tables that you want to view.")
	cleanrooms_listIdMappingTablesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listIdMappingTablesCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_listIdMappingTablesCmd)
}
