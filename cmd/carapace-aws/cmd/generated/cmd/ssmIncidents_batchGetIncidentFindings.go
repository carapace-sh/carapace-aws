package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_batchGetIncidentFindingsCmd = &cobra.Command{
	Use:   "batch-get-incident-findings",
	Short: "Retrieves details about all specified findings for an incident, including descriptive details about each finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_batchGetIncidentFindingsCmd).Standalone()

	ssmIncidents_batchGetIncidentFindingsCmd.Flags().String("finding-ids", "", "A list of IDs of findings for which you want to view details.")
	ssmIncidents_batchGetIncidentFindingsCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident for which you want to view finding details.")
	ssmIncidents_batchGetIncidentFindingsCmd.MarkFlagRequired("finding-ids")
	ssmIncidents_batchGetIncidentFindingsCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_batchGetIncidentFindingsCmd)
}
