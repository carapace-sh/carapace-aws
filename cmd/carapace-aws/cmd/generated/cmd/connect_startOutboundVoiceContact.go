package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startOutboundVoiceContactCmd = &cobra.Command{
	Use:   "start-outbound-voice-contact",
	Short: "Places an outbound call to a contact, and then initiates the flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startOutboundVoiceContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_startOutboundVoiceContactCmd).Standalone()

		connect_startOutboundVoiceContactCmd.Flags().String("answer-machine-detection-config", "", "Configuration of the answering machine detection for this outbound call.")
		connect_startOutboundVoiceContactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
		connect_startOutboundVoiceContactCmd.Flags().String("campaign-id", "", "The campaign identifier of the outbound communication.")
		connect_startOutboundVoiceContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_startOutboundVoiceContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow for the outbound call.")
		connect_startOutboundVoiceContactCmd.Flags().String("description", "", "A description of the voice contact that appears in the agent's snapshot in the CCP logs.")
		connect_startOutboundVoiceContactCmd.Flags().String("destination-phone-number", "", "The phone number of the customer, in E.164 format.")
		connect_startOutboundVoiceContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_startOutboundVoiceContactCmd.Flags().String("name", "", "The name of a voice contact that is shown to an agent in the Contact Control Panel (CCP).")
		connect_startOutboundVoiceContactCmd.Flags().String("outbound-strategy", "", "Information about the outbound strategy.")
		connect_startOutboundVoiceContactCmd.Flags().String("queue-id", "", "The queue for the call.")
		connect_startOutboundVoiceContactCmd.Flags().String("references", "", "A formatted URL that is shown to an agent in the Contact Control Panel (CCP).")
		connect_startOutboundVoiceContactCmd.Flags().String("related-contact-id", "", "The `contactId` that is related to this contact.")
		connect_startOutboundVoiceContactCmd.Flags().String("source-phone-number", "", "The phone number associated with the Amazon Connect instance, in E.164 format.")
		connect_startOutboundVoiceContactCmd.Flags().String("traffic-type", "", "Denotes the class of traffic.")
		connect_startOutboundVoiceContactCmd.MarkFlagRequired("contact-flow-id")
		connect_startOutboundVoiceContactCmd.MarkFlagRequired("destination-phone-number")
		connect_startOutboundVoiceContactCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_startOutboundVoiceContactCmd)
}
