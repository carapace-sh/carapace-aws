package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_updateResponsePlanCmd = &cobra.Command{
	Use:   "update-response-plan",
	Short: "Updates the specified response plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_updateResponsePlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_updateResponsePlanCmd).Standalone()

		ssmIncidents_updateResponsePlanCmd.Flags().String("actions", "", "The actions that this response plan takes at the beginning of an incident.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the response plan.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("chat-channel", "", "The Chatbot chat channel used for collaboration during an incident.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("client-token", "", "A token ensuring that the operation is called only once with the specified details.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("display-name", "", "The long format name of the response plan.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("engagements", "", "The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("incident-template-dedupe-string", "", "The string Incident Manager uses to prevent duplicate incidents from being created by the same incident in the same account.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("incident-template-impact", "", "Defines the impact to the customers.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("incident-template-notification-targets", "", "The Amazon SNS targets that are notified when updates are made to an incident.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("incident-template-summary", "", "A brief summary of the incident.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("incident-template-tags", "", "Tags to assign to the template.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("incident-template-title", "", "The short format name of the incident.")
		ssmIncidents_updateResponsePlanCmd.Flags().String("integrations", "", "Information about third-party services integrated into the response plan.")
		ssmIncidents_updateResponsePlanCmd.MarkFlagRequired("arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_updateResponsePlanCmd)
}
