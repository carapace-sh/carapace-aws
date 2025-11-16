package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_populateIdMappingTableCmd = &cobra.Command{
	Use:   "populate-id-mapping-table",
	Short: "Defines the information that's necessary to populate an ID mapping table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_populateIdMappingTableCmd).Standalone()

	cleanrooms_populateIdMappingTableCmd.Flags().String("id-mapping-table-identifier", "", "The unique identifier of the ID mapping table that you want to populate.")
	cleanrooms_populateIdMappingTableCmd.Flags().String("job-type", "", "The job type of the rule-based ID mapping job.")
	cleanrooms_populateIdMappingTableCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID mapping table that you want to populate.")
	cleanrooms_populateIdMappingTableCmd.MarkFlagRequired("id-mapping-table-identifier")
	cleanrooms_populateIdMappingTableCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_populateIdMappingTableCmd)
}
