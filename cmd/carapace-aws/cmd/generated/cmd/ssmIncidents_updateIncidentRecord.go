package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_updateIncidentRecordCmd = &cobra.Command{
	Use:   "update-incident-record",
	Short: "Update the details of an incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_updateIncidentRecordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_updateIncidentRecordCmd).Standalone()

		ssmIncidents_updateIncidentRecordCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the incident record you are updating.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("chat-channel", "", "The Chatbot chat channel where responders can collaborate.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("client-token", "", "A token that ensures that a client calls the operation only once with the specified details.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("impact", "", "Defines the impact of the incident to customers and applications.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("notification-targets", "", "The Amazon SNS targets that Incident Manager notifies when a client updates an incident.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("status", "", "The status of the incident.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("summary", "", "A longer description of what occurred during the incident.")
		ssmIncidents_updateIncidentRecordCmd.Flags().String("title", "", "A brief description of the incident.")
		ssmIncidents_updateIncidentRecordCmd.MarkFlagRequired("arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_updateIncidentRecordCmd)
}
