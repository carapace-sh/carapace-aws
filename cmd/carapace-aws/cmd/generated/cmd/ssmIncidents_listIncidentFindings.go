package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listIncidentFindingsCmd = &cobra.Command{
	Use:   "list-incident-findings",
	Short: "Retrieves a list of the IDs of findings, plus their last modified times, that have been identified for a specified incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listIncidentFindingsCmd).Standalone()

	ssmIncidents_listIncidentFindingsCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident for which you want to view associated findings.")
	ssmIncidents_listIncidentFindingsCmd.Flags().String("max-results", "", "The maximum number of findings to retrieve per call.")
	ssmIncidents_listIncidentFindingsCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
	ssmIncidents_listIncidentFindingsCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_listIncidentFindingsCmd)
}
