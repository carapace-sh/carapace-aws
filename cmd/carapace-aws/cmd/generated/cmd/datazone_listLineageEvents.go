package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listLineageEventsCmd = &cobra.Command{
	Use:   "list-lineage-events",
	Short: "Lists lineage events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listLineageEventsCmd).Standalone()

	datazone_listLineageEventsCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list lineage events.")
	datazone_listLineageEventsCmd.Flags().String("max-results", "", "The maximum number of lineage events to return in a single call to ListLineageEvents.")
	datazone_listLineageEventsCmd.Flags().String("next-token", "", "When the number of lineage events is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of lineage events, the response includes a pagination token named NextToken.")
	datazone_listLineageEventsCmd.Flags().String("processing-status", "", "The processing status of a lineage event.")
	datazone_listLineageEventsCmd.Flags().String("sort-order", "", "The sort order of the lineage events.")
	datazone_listLineageEventsCmd.Flags().String("timestamp-after", "", "The after timestamp of a lineage event.")
	datazone_listLineageEventsCmd.Flags().String("timestamp-before", "", "The before timestamp of a lineage event.")
	datazone_listLineageEventsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listLineageEventsCmd)
}
