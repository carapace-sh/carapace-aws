package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_startIncidentCmd = &cobra.Command{
	Use:   "start-incident",
	Short: "Used to start an incident from CloudWatch alarms, EventBridge events, or manually.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_startIncidentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_startIncidentCmd).Standalone()

		ssmIncidents_startIncidentCmd.Flags().String("client-token", "", "A token ensuring that the operation is called only once with the specified details.")
		ssmIncidents_startIncidentCmd.Flags().String("impact", "", "Defines the impact to the customers.")
		ssmIncidents_startIncidentCmd.Flags().String("related-items", "", "Add related items to the incident for other responders to use.")
		ssmIncidents_startIncidentCmd.Flags().String("response-plan-arn", "", "The Amazon Resource Name (ARN) of the response plan that pre-defines summary, chat channels, Amazon SNS topics, runbooks, title, and impact of the incident.")
		ssmIncidents_startIncidentCmd.Flags().String("title", "", "Provide a title for the incident.")
		ssmIncidents_startIncidentCmd.Flags().String("trigger-details", "", "Details of what created the incident record in Incident Manager.")
		ssmIncidents_startIncidentCmd.MarkFlagRequired("response-plan-arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_startIncidentCmd)
}
