package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_deleteIncidentRecordCmd = &cobra.Command{
	Use:   "delete-incident-record",
	Short: "Delete an incident record from Incident Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_deleteIncidentRecordCmd).Standalone()

	ssmIncidents_deleteIncidentRecordCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the incident record you are deleting.")
	ssmIncidents_deleteIncidentRecordCmd.MarkFlagRequired("arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_deleteIncidentRecordCmd)
}
