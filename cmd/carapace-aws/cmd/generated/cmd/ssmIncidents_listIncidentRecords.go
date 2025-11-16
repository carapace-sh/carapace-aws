package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listIncidentRecordsCmd = &cobra.Command{
	Use:   "list-incident-records",
	Short: "Lists all incident records in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listIncidentRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_listIncidentRecordsCmd).Standalone()

		ssmIncidents_listIncidentRecordsCmd.Flags().String("filters", "", "Filters the list of incident records you want to search through.")
		ssmIncidents_listIncidentRecordsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
		ssmIncidents_listIncidentRecordsCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_listIncidentRecordsCmd)
}
