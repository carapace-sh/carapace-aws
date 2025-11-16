package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_createResponsePlanCmd = &cobra.Command{
	Use:   "create-response-plan",
	Short: "Creates a response plan that automates the initial response to incidents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_createResponsePlanCmd).Standalone()

	ssmIncidents_createResponsePlanCmd.Flags().String("actions", "", "The actions that the response plan starts at the beginning of an incident.")
	ssmIncidents_createResponsePlanCmd.Flags().String("chat-channel", "", "The Chatbot chat channel used for collaboration during an incident.")
	ssmIncidents_createResponsePlanCmd.Flags().String("client-token", "", "A token ensuring that the operation is called only once with the specified details.")
	ssmIncidents_createResponsePlanCmd.Flags().String("display-name", "", "The long format of the response plan name.")
	ssmIncidents_createResponsePlanCmd.Flags().String("engagements", "", "The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.")
	ssmIncidents_createResponsePlanCmd.Flags().String("incident-template", "", "Details used to create an incident when using this response plan.")
	ssmIncidents_createResponsePlanCmd.Flags().String("integrations", "", "Information about third-party services integrated into the response plan.")
	ssmIncidents_createResponsePlanCmd.Flags().String("name", "", "The short format name of the response plan.")
	ssmIncidents_createResponsePlanCmd.Flags().String("tags", "", "A list of tags that you are adding to the response plan.")
	ssmIncidents_createResponsePlanCmd.MarkFlagRequired("incident-template")
	ssmIncidents_createResponsePlanCmd.MarkFlagRequired("name")
	ssmIncidentsCmd.AddCommand(ssmIncidents_createResponsePlanCmd)
}
