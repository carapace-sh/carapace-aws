package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_getIncidentRecordCmd = &cobra.Command{
	Use:   "get-incident-record",
	Short: "Returns the details for the specified incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_getIncidentRecordCmd).Standalone()

	ssmIncidents_getIncidentRecordCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the incident record.")
	ssmIncidents_getIncidentRecordCmd.MarkFlagRequired("arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_getIncidentRecordCmd)
}
