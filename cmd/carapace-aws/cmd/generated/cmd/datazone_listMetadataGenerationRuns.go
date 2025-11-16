package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listMetadataGenerationRunsCmd = &cobra.Command{
	Use:   "list-metadata-generation-runs",
	Short: "Lists all metadata generation runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listMetadataGenerationRunsCmd).Standalone()

	datazone_listMetadataGenerationRunsCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where you want to list metadata generation runs.")
	datazone_listMetadataGenerationRunsCmd.Flags().String("max-results", "", "The maximum number of metadata generation runs to return in a single call to ListMetadataGenerationRuns.")
	datazone_listMetadataGenerationRunsCmd.Flags().String("next-token", "", "When the number of metadata generation runs is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of metadata generation runs, the response includes a pagination token named NextToken.")
	datazone_listMetadataGenerationRunsCmd.Flags().String("status", "", "The status of the metadata generation runs.")
	datazone_listMetadataGenerationRunsCmd.Flags().String("type", "", "The type of the metadata generation runs.")
	datazone_listMetadataGenerationRunsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listMetadataGenerationRunsCmd)
}
